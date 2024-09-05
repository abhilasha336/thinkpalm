package utilities

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/IBM/sarama"
)

var kafkaBrokers = []string{"localhost:9092"}

type ValidationResponse struct {
	UserName string `json:"username"`

	ClientId string                 `json:"clientid"`
	Data     map[string]interface{} `json:",omitempty"`
}

// Generate random client credentials
func GenerateClientCredentials() (clientID, clientSecret string, err error) {
	clientID = "client-id-" + GenerateRandomString(8)
	clientSecretBytes := make([]byte, 32)
	_, err = rand.Read(clientSecretBytes)
	if err != nil {
		return "", "", err
	}
	clientSecret = fmt.Sprintf("%x", clientSecretBytes)
	return
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GetKafkaProducer() sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(kafkaBrokers, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	return producer
}

func PublishValidationResult(response ValidationResponse) {
	producer := GetKafkaProducer()
	defer producer.Close()

	responseMessage, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return
	}

	message := &sarama.ProducerMessage{
		Topic: "metadata",
		Value: sarama.StringEncoder(responseMessage),
	}

	if _, _, err := producer.SendMessage(message); err != nil {
		log.Printf("Error writing message to Kafka: %v", err)
	}
}

func GetKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(kafkaBrokers, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	return consumer
}
