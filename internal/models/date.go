package models

type DateResponse struct {
	Dates []string `json:"dates"`
}

type Date struct {
	DateTime string `json:"datetime"`
}
