package main

import (
	"fmt"
	"log"
)

// An entire forcast/list of weather conditions (independent of api)
type Weather struct {
	Sunrise int
	Sunset  int
	Current Conditions
	Hourly  []Conditions
}

// Weather conditions for a specific point in time (independent of api)
type Conditions struct {
	Clouds      int
	Humidity    int
	Temperature float64
	FeelsLike   float64
	WindSpeed   float64
	WindDeg     int
	Icon        string
	Description string
}

func GetWeather(location *UserLocation) (Weather, error) {
	response, err := GetOpenWeatherConditions(location)

	if err != nil {
		log.Print("Failed to get weather\n", err)
		return Weather{}, nil
	}

	return response.normalize(), nil
}

// JSON Response from OpenWeatherMap
type OpenWeatherConditionsResponse struct {
	Current struct {
		Clouds      int     `json:"clouds"`
		Humidity    int     `json:"humidity"`
		Sunrise     int     `json:"sunrise"`
		Sunset      int     `json:"sunset"`
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		WindSpeed   float64 `json:"wind_speed"`
		WindDeg     int     `json:"wind_deg"`

		Weather []struct {
			Icon        string `json:"icon"`
			Description string `json:"description"`
		}
	}
}

// Convert OpenWeatherMap api response to a universal Weather struct
func (response *OpenWeatherConditionsResponse) normalize() Weather {
	weather := Weather{
		Sunrise: response.Current.Sunrise,
		Sunset:  response.Current.Sunset,
		Current: Conditions{
			Clouds:      response.Current.Clouds,
			Humidity:    response.Current.Humidity,
			Temperature: response.Current.Temperature,
			FeelsLike:   response.Current.FeelsLike,
			WindSpeed:   response.Current.WindSpeed,
			WindDeg:     response.Current.WindDeg,
			Icon:        response.Current.Weather[0].Icon,
			Description: response.Current.Weather[0].Description,
		},
	}

	return weather
}

const exclude string = "minutely,daily,alerts"
const units string = "imperial"

func GetOpenWeatherConditions(location *UserLocation) (OpenWeatherConditionsResponse, error) {

	current := OpenWeatherConditionsResponse{}
	err := GetJsonInStruct(
		fmt.Sprintf(
			"%s?lat=%f&lon=%f&exclude=%s&units=%s&appid=%s",
			env.GetString("OPENWEATHERMAP_URL"),
			location.Latitude,
			location.Longitude,
			exclude,
			units,
			env.GetString("OPENWEATHERMAP_KEY"),
		),
		&current,
	)

	if err != nil {
		log.Print("Failed to obtain weather conditions\n", err)
		return OpenWeatherConditionsResponse{}, err
	}

	return current, nil
}
