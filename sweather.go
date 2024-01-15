package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port = 8000

var env = EnvironmentVariables{FilePath: "./.env"}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprintf(w, "An error occurred while generating your page!%s", err)
		return
	}

	ipaddy := ReadUserIP(r)
	location, err := RequestLocation(ipaddy)
	conditions, err := RequestCurrentConditions(&location)

	template.Execute(w, Alaska)
	log.Printf("Got lat, long: %f, %f", location.latitude, location.longitude)
	log.Printf("Got conditions: %s", conditions)
}

func main() {
	env.Load()

	http.HandleFunc("/", homePageHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
