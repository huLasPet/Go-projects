package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/joho/godotenv"
	//    "github.com/aws/aws-sdk-go/service/s3"
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

func getVoices(pollySession *session.Session) {
	svc := polly.New(pollySession)
	input := &polly.DescribeVoicesInput{LanguageCode: aws.String("en-US")}
	resp, _ := svc.DescribeVoices(input)
	for _, v := range resp.Voices {
		fmt.Println("Name:   " + *v.Name)
		fmt.Println("Gender: " + *v.Gender)
		fmt.Println("")
	}
}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pollySession := createSession()
	getVoices(pollySession)

}
