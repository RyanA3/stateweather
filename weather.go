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
	DateTime    int
	UVIndex     float64
	Visibility  int
	Clouds      int
	Humidity    int
	DewPoint    float64
	Temperature float64
	FeelsLike   float64
	WindSpeed   float64
	WindGust    float64
	WindDeg     int
	Pressure    int
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
type OpenWeatherResponse struct {
	Current OpenWeatherConditions
}

type OpenWeatherConditions struct {
	DateTime    int     `json:"dt"`
	Visibility  int     `json:"visibility"`
	UVIndex     float64 `json:"uvi"`
	Pressure    int     `json:"pressure"`
	DewPoint    float64 `json:"dew_point"`
	Clouds      int     `json:"clouds"`
	Humidity    int     `json:"humidity"`
	Sunrise     int     `json:"sunrise"`
	Sunset      int     `json:"sunset"`
	Temperature float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_like"`
	WindSpeed   float64 `json:"wind_speed"`
	WindGust    float64 `json:"wind_gust"`
	WindDeg     int     `json:"wind_deg"`

	Weather []struct {
		Icon        string `json:"icon"`
		Description string `json:"description"`
	}
}

// Convert an instance of weather conditions from OpenWeatherMap api response to an instance of the generic weather conditions struct
func (responseConditions *OpenWeatherConditions) normalize() Conditions {
	return Conditions{
		DateTime:    responseConditions.DateTime,
		Visibility:  responseConditions.Visibility,
		UVIndex:     responseConditions.UVIndex,
		Pressure:    responseConditions.Pressure,
		DewPoint:    responseConditions.DewPoint,
		WindGust:    responseConditions.WindGust,
		Clouds:      responseConditions.Clouds,
		Humidity:    responseConditions.Humidity,
		Temperature: responseConditions.Temperature,
		FeelsLike:   responseConditions.FeelsLike,
		WindSpeed:   responseConditions.WindSpeed,
		WindDeg:     responseConditions.WindDeg,
		Icon:        responseConditions.Weather[0].Icon,
		Description: responseConditions.Weather[0].Description,
	}
}

// Convert OpenWeatherMap api response to a universal Weather struct
func (response *OpenWeatherResponse) normalize() Weather {
	weather := Weather{
		Sunrise: response.Current.Sunrise,
		Sunset:  response.Current.Sunset,
		Current: response.Current.normalize(),
	}

	return weather
}

const exclude string = "minutely,daily,alerts"
const units string = "imperial"

func GetOpenWeatherConditions(location *UserLocation) (OpenWeatherResponse, error) {

	current := OpenWeatherResponse{}
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
		return OpenWeatherResponse{}, err
	}

	return current, nil
}
