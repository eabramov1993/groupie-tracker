package models

type LocationsRespons struct {
	Locations []string `json:"locations"`
}

type Location struct {
	City string `json:"city"`
}
