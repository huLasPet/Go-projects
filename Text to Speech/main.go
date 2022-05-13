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

func renderTemplate(w http.ResponseWriter, page string) {
	templateVariable, err := template.ParseFiles(page)
	if err != nil {
		panic(err)
	}
	err = templateVariable.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func homePage(responseWriter http.ResponseWriter, r *http.Request) {
	renderTemplate(responseWriter, "Files/index.html")

}

func startSynth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		tts := r.Form["tts"][0]
		voice := r.Form["voice"][0]
		pollySession := AWSPolly.CreateSession()
		AWSPolly.SynthSpeach(pollySession, tts, voice)
	}

}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/startsynth", startSynth)
	http.ListenAndServe(":8080", nil)
}
