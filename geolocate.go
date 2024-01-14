package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

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
