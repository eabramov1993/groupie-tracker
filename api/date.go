package api

import (
	"encoding/json"
	"groupie-trackers/internal/models"
	"net/http"
)

func GetDate() ([]models.Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var Date []models.Date

	err = json.NewDecoder(resp.Body).Decode(&Date)
	if err != nil {
		return nil, err
	}

	return Date, nil
}
