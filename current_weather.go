package main

import (
	"fmt"
	"log"
)

// This is fucked
type Conditions struct {

	// What is this
	Current struct {
		Clouds      int     `json:"clouds"`
		Humidity    int     `json:"humidity"`
		Sunrise     int     `json:"sunrise"`
		Sunset      int     `json:"sunset"`
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		WindSpeed   float64 `json:"wind_speed"`
		WindDeg     int     `json:"wind_deg"`

		//Why must we nest structs to unmarshal json
		Weather []struct {
			Icon        string `json:"icon"`
			Description string `json:"description"`

			//I'm crying
			//This can't even be used in an html template because it's an array
		}
	}
}

const exclude string = "minutely,hourly,daily,alerts"
const units string = "imperial"

func RequestCurrentConditions(location *UserLocation) (Conditions, error) {

	current := Conditions{}
	err := GetJsonInStruct(
		fmt.Sprintf(
			"%s?lat=%f&lon=%f&exclude=%s&units=%s&appid=%s",
			env.GetString("OPENWEATHERMAP_URL"),
			location.latitude,
			location.longitude,
			exclude,
			units,
			env.GetString("OPENWEATHERMAP_KEY"),
		),
		&current,
	)

	if err != nil {
		log.Print("Failed to obtain weather conditions\n", err)
		return Conditions{}, err
	}

	return current, nil
}
