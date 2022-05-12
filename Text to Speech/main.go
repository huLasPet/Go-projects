package main

import (
	"fmt"
	"io"
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

func synthSpeach(pollySession *session.Session) {
	svc := polly.New(pollySession)
	input := &polly.SynthesizeSpeechInput{OutputFormat: aws.String("mp3"), Text: aws.String("Some text here to read"), VoiceId: aws.String("Joanna")}
	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		fmt.Println("Got error calling SynthesizeSpeech:")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	mp3File := "test.mp3"
	outFile, _ := os.Create(mp3File)
	defer outFile.Close()

	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		fmt.Println("Got error saving MP3:")
		fmt.Print(err.Error())
		os.Exit(1)
	}

}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pollySession := createSession()
	//getVoices(pollySession)
	synthSpeach(pollySession)

}
