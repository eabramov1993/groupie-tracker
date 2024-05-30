package api

import (
	"encoding/json"
	"fmt"
	"groupie-trackers/internal/models"
	"net/http"
)

func GetDates(artistID int) ([]models.Date, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", artistID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get dates: %v", err)
	}

	defer resp.Body.Close()

	var datesRespons models.DateResponse

	err = json.NewDecoder(resp.Body).Decode(&datesRespons)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates: %v", err)
	}

	var dates []models.Date
	for _, dataTime := range datesRespons.Dates {
		dates = append(dates, models.Date{DateTime: dataTime})
	}
	return dates, nil
}
