package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type UserLocation struct {
	latitude           float64
	longitude          float64
	city_geoname_id    string
	country_geoname_id string
}

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
func RequestLocation(ipaddy string) (UserLocation, error) {

	response, err := GetJson(
		fmt.Sprintf("%s?api_key=%s&ip_address=%s",
			env.GetString("ABSTRACT_GEOLOCATION_URL"),
			env.GetString("ABSTRACT_GEOLOCATION_KEY"),
			ipaddy,
		),
	)

	if err != nil {
		log.Print("Failed to obtain user's location\n ", err)
		return UserLocation{}, err
	}

	fmt.Println(response)

	location := UserLocation{
		latitude:           response["latitude"].(float64),
		longitude:          response["longitude"].(float64),
		city_geoname_id:    response["city_geoname_id"].(string),
		country_geoname_id: response["country_geoname_id"].(string),
	}

	return location, nil
}
