package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	tokenValidationURL = "http://localhost:8080/validate" // Replace with your actual token validation URL
)

// TokenValidationMiddleware is a Gin middleware that validates the token
func TokenValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.RequestURI()
		fmt.Println("haaaaaaaaaaaaaaaaaaaaaaaaaaa", path)
		sli := strings.Split(path, "?")
		fmt.Println("sli++++++++++++++++++++++++", sli)
		path = sli[0]

		if path == "/client-register-form" {
			// Get token from Authorization header
			token := c.Query("Authorization")
			if token == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
				c.Abort()
				return
			}
			fmt.Println("tokkkk----------------------------------++++++", token)

			// // Call token validation endpoint
			// resp, err := http.Post(tokenValidationURL, "application/json", bytes.NewBufferString(`{"token":"`+strings.TrimPrefix(token, "Bearer ")+`"}`))
			// if err != nil || resp.StatusCode != http.StatusOK {
			// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			// 	c.Abort()
			// 	return
			// }
			token = strings.TrimPrefix(token, "Bearer ")
			jsonBody := fmt.Sprintf(`{"token":"%s"}`, token)
			resp, err := http.Post(tokenValidationURL, "application/json", bytes.NewBufferString(jsonBody))
			if err != nil || resp.StatusCode != http.StatusOK {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}

			fmt.Println("s-----------------------------------------------------", resp.StatusCode)

			// Parse the validation response
			var validationResponse map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&validationResponse); err != nil || !validationResponse["valid"].(bool) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}

			c.Next()

		} else {
			c.Next()
		}

	}
}
