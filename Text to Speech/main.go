package main

import (
	"AWSPolly/AWSPolly"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

const envFile = "/Users/nbyy/Library/CloudStorage/OneDrive-Personal/Golang Round 1/Golang env files/.env"

func createSession() *session.Session {
	ttsKey := os.Getenv("tts_key")
	ttsSecret := os.Getenv("tts_secret")

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(ttsKey, ttsSecret, ""),
	})
	return sess
}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pollySession := createSession()
	AWSPolly.GetVoices(pollySession)
	AWSPolly.SynthSpeach(pollySession)

	//TODO: Create a webapp from this
}
