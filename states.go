package main

import (
	"math"
	"time"
)

type State struct {
	Name             string
	ImageName        string
	ImageLicense     string
	ImageLicenseLink string

	WeatherAverages map[Season]WeatherAverages
}

type WeatherAverages struct {
	Temperature float64
	Sunshine    int
	Humidity    int
}

type Season int

const (
	Summer = 0 + iota
	Autumn
	Winter
	Spring
)

func GetSeason(datetime time.Time) Season {
	switch datetime.Month() {
	case time.March, time.April, time.May:
		return Spring
	case time.June, time.July, time.August:
		return Summer
	case time.September, time.October, time.November:
		return Autumn
	default:
		return Winter
	}
}

const tempRange float64 = 100

func (state *State) CompareWeather(conditions Conditions, season Season) float64 {
	averages := state.WeatherAverages[season]

	dtemp := math.Max(math.Abs(conditions.Temperature-averages.Temperature), tempRange)
	ntemp := dtemp / tempRange

	//Results will range from 0 to 1 (similarity)

	return ntemp
}

var States = [...]State{
	{Name: "Alaska", ImageName: "alaska.jpg"},
}
