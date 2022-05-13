package main

import (
	"AWSPolly/AWSPolly"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/joho/godotenv"
)

const envFile = "/Users/nbyy/Library/CloudStorage/OneDrive-Personal/Golang Round 1/Golang env files/.env"

func startSynth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Files/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		tts := r.Form["tts"][0]
		voice := r.Form["voice"][0]
		pollySession := AWSPolly.CreateSession()
		AWSPolly.SynthSpeach(pollySession, tts, voice)
		fmt.Fprintf(w, "Done")
	}
}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/", startSynth)
	http.ListenAndServe(":8080", nil)
}
