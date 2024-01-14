package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

const port = 8000

var env = EnvironmentVariables{FilePath: "./.env"}

type State struct {
	Name      string
	ImageName string
}

var Alaska State = State{Name: "Alaska", ImageName: "alaska.jpg"}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	IPAddress = strings.Split(IPAddress, ":")[0]

	return IPAddress
}

// Makes a request to the Abstract geolocation api to get the user's latitude and longitude
func RequestLatLong(ipaddy string) (float64, float64) {

	response, err := http.Get(
		fmt.Sprintf("%s?api_key=%s&ip_address=%s",
			env.GetString("ABSTRACT_GEOLOCATION_URL"),
			env.GetString("ABSTRACT_GEOLOCATION_KEY"),
			ipaddy,
		),
	)

	if err != nil {
		log.Print("Failed to obtain user's lat/long\n ", err)
		return 0, 0
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Print("Failed to read geolocation request response\n", err)
		return 0, 0
	}

	log.Println(string(body))
	return 0, 0
}

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
