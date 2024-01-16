package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type UserLocation struct {
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	CityGeonameId    float64 `json:"city_geoname_id"`
	CountryGeonameId float64 `json:"country_geoname_id"`
	TimeZone         struct {
		Name      string `json:"name"`
		GMTOffset int    `json:"gmt_offset"`
	}
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
	location := UserLocation{}

	err := GetJsonInStruct(
		fmt.Sprintf("%s?api_key=%s&ip_address=%s",
			env.GetString("ABSTRACT_GEOLOCATION_URL"),
			env.GetString("ABSTRACT_GEOLOCATION_KEY"),
			ipaddy,
		),
		&location,
	)

	if err != nil {
		log.Print("Failed to obtain user's location\n ", err)
		return UserLocation{}, err
	}

	return location, nil
}
