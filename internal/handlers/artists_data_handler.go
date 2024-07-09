package handlers

import (
	"encoding/json"
	"fmt" // Для форматирования сообщения об ошибке
	"groupie-trackers/api"
	"net/http"
)

func ArtistsDataHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что используется метод GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	artists, err := api.GetArtists()
	if err != nil {
		// Обработка ошибки при получении данных об артистах
		errorMessage := fmt.Sprintf("Error getting artists: %v", err)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type для JSON
	w.Header().Set("Content-Type", "application/json")

	// Кодируем данные в JSON и отправляем
	err = json.NewEncoder(w).Encode(artists)
	if err != nil {
		// Обработка ошибки при кодировании JSON
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
