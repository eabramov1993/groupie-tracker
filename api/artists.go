package api

import (
	"encoding/json"
	"groupie-trackers/internal/models"
	"net/http"
)

func GetArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var Artists []models.Artist

	err = json.NewDecoder(resp.Body).Decode(&Artists)
	if err != nil {
		return nil, err
	}

	return Artists, nil

}
