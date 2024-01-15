package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port = 8000

var env = EnvironmentVariables{FilePath: "./.env"}

type PageContent struct {
	State      State
	Conditions Conditions
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprintf(w, "An error occurred while generating your page!%s", err)
		return
	}

	ipaddy := ReadUserIP(r)
	location, err := RequestLocation(ipaddy)
	conditions, err := RequestCurrentConditions(&location)

	content := PageContent{State: Alaska, Conditions: conditions}

	template.Execute(w, content)
	log.Printf("Got lat, long: %f, %f", location.Latitude, location.Longitude)
	log.Printf("Got conditions: %s, temp: %f", conditions.Current.Weather[0].Description, conditions.Current.Temperature)
}

func main() {
	env.Load()

	http.HandleFunc("/", homePageHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
