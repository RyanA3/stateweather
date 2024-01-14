package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port = 8000

var env = EnvironmentVariables{FilePath: "./.env"}

type State struct {
	Name      string
	ImageName string
}

var Alaska State = State{Name: "Alaska", ImageName: "alaska.jpg"}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprintf(w, "An error occurred while generating your page!%s", err)
		return
	}

	ipaddy := ReadUserIP(r)
	lat, long := RequestLatLong(ipaddy)

	template.Execute(w, Alaska)

	fmt.Fprintf(w, ipaddy)

	log.Printf("Got lat, long: %f, %f", lat, long)
}

func main() {
	env.Load()

	http.HandleFunc("/", homePageHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
