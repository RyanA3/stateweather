package main

type State struct {
	Name             string
	ImageName        string
	ImageLicense     string
	ImageLicenseLink string
}

var Alaska State = State{Name: "Alaska", ImageName: "alaska.jpg"}
