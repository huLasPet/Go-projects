package main

import (
	"AWSPolly/AWSPolly"
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
		http.Redirect(w, r, "/play", http.StatusSeeOther)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Files/synthedText.mp3")
}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/", startSynth)
	http.HandleFunc("/play", play)
	http.ListenAndServe("192.168.0.27:8080", nil)
}
