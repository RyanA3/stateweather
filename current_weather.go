package main

import (
	"fmt"
	"log"
)

type Conditions struct {
}

const exclude string = "minutely,hourly,daily,alerts"

func RequestCurrentConditions(location *UserLocation) (Conditions, error) {

	response, err := GetJson(
		fmt.Sprintf(
			"%s?lat=%f&lon=%f&exclude=%s&appid=%s",
			env.GetString("OPENWEATHERMAP_URL"),
			location.latitude,
			location.longitude,
			exclude,
			env.GetString("OPENWEATHERMAP_KEY"),
		),
	)

	if err != nil {
		log.Print("Failed to obtain weather conditions\n", err)
		return Conditions{}, err
	}

	log.Print(response)
	return Conditions{}, nil
}
