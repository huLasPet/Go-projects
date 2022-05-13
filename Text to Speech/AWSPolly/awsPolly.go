package AWSPolly

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

func GetVoices(pollySession *session.Session) {
	svc := polly.New(pollySession)
	input := &polly.DescribeVoicesInput{LanguageCode: aws.String("en-US")}
	resp, _ := svc.DescribeVoices(input)
	for _, v := range resp.Voices {
		fmt.Println("Name:   " + *v.Name)
		fmt.Println("Gender: " + *v.Gender)
		fmt.Println("")
	}
}

func SynthSpeach(pollySession *session.Session, text, voice, filename string) {
	svc := polly.New(pollySession)
	input := &polly.SynthesizeSpeechInput{OutputFormat: aws.String("mp3"), Text: aws.String(text), VoiceId: aws.String(voice), Engine: aws.String("neural")}
	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		fmt.Println("Got error calling SynthesizeSpeech:")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//mp3File := "pollySynth.mp3"
	outFile, _ := os.Create(filename + "/synthedText.mp3")
	defer outFile.Close()

	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		fmt.Println("Got error saving MP3:")
		fmt.Print(err.Error())
		os.Exit(1)
	}

}

func CreateSession() *session.Session {
	ttsKey := os.Getenv("tts_key")
	ttsSecret := os.Getenv("tts_secret")

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(ttsKey, ttsSecret, ""),
	})
	return sess
}
