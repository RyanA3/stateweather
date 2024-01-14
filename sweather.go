package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port = 8000

type State struct {
	Name  string
	Image string
}

var Alaska State = State{Name: "Alaska", Image: "alaska.jpg"}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// func GetLatLong(ipaddy string) (float64, float64) {

// }

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprintf(w, "An error occurred while generating your page!%s", err)
		return
	}

	template.Execute(w, Alaska)

	ipaddy := ReadUserIP(r)
	fmt.Fprintf(w, ipaddy)

}

func main() {
	env := EnvironmentVariables{FilePath: "./.env"}
	env.Load()

	log.Printf("KABOOM=%t", env.GetBoolean("KABOOM"))

	http.HandleFunc("/", homePageHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
