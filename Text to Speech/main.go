package main

import (
	"AWSPolly/AWSPolly"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/sqweek/dialog"
)

const envFile = "/Users/nbyy/Library/CloudStorage/OneDrive-Personal/Golang Round 1/Golang env files/.env"

func startSynth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Files/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		filename, err := dialog.Directory().Title("Save to:").Browse()
		if err != nil {
			filename = "Files/"
		}
		tts := r.Form["tts"][0]
		voice := r.Form["voice"][0]
		pollySession := AWSPolly.CreateSession()
		AWSPolly.SynthSpeach(pollySession, tts, voice, filename)
		fmt.Fprintf(w, "Done")
	}
}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/", startSynth)
	http.ListenAndServe("192.168.0.27:8080", nil)
}
