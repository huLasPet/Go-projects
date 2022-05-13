package main

import (
	"AWSPolly/AWSPolly"
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
	pollySession := AWSPolly.CreateSession()
	AWSPolly.SynthSpeach(pollySession)
}

func main() {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", startSynth)
	http.ListenAndServe(":8080", nil)

	//AWSPolly.GetVoices(pollySession)

	//TODO: Create a webapp from this
}
