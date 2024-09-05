package controllers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/abhilasha336/thinkpalm/internal/utilities"
	"github.com/gin-gonic/gin"
)

func (think *ThinkpalmController) ClientRegisterFormSubmit(ctx *gin.Context) {
	// Parse multipart form
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	// Retrieve form fields
	formData := form.Value
	relaystate := formData["relaystate"][0]
	entityid := formData["entityid"][0]
	recipient := formData["recipient"][0]
	counsumervalidatorurl := formData["counsumervalidatorurl"][0]
	slo := formData["slo"][0]
	sso := formData["sso"][0]
	samlnamedid := formData["samlnamedid"][0]

	// Retrieve PEM file
	file, _, err := ctx.Request.FormFile("samlcertificate")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}
	defer file.Close()

	// Read file content into a string
	var fileContent bytes.Buffer
	_, err = io.Copy(&fileContent, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	samlCertificate := fileContent.String()

	// Generate client credentials
	clientID, clientSecret, err := utilities.GenerateClientCredentials()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate client credentials"})
		return
	}

	// Populate map with all data
	result := map[string]interface{}{
		"relaystate":            relaystate,
		"entityid":              entityid,
		"recipient":             recipient,
		"counsumervalidatorurl": counsumervalidatorurl,
		"slo":                   slo,
		"sso":                   sso,
		"samlnamedid":           samlnamedid,
		"samlcertificate":       samlCertificate,
		"client_id":             clientID,
		"client_secret":         clientSecret,
	}

	// Publish to Kafka
	username := result["samlnamedid"].(string)
	if err := think.useCase.InsertClientConfig(result, clientID, clientSecret); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to store app config in db"})
		return
	}
	var res utilities.ValidationResponse
	res.UserName = username
	res.ClientId = clientID
	res.Data = result
	utilities.PublishValidationResult(res)

	// Respond to the client
	ctx.JSON(http.StatusOK, gin.H{
		"clientid":     clientID,
		"clientsecret": clientSecret,
	})
}
