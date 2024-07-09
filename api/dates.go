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

	var datesResponse models.DatesResponse
	err = json.NewDecoder(resp.Body).Decode(&datesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates: %v", err)
	}

	var dates []models.Date
	for _, dateTime := range datesResponse.Dates {
		dates = append(dates, models.Date{DateTime: dateTime})
	}

	return dates, nil
}
