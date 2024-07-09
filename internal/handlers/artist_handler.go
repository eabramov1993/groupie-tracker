package handlers

import (
	"groupie-trackers/api"
	"groupie-trackers/internal/models"
	"net/http"
	"strconv"
	"text/template"
)

// ArtistHandler обрабатывает запросы к странице конкретного артиста, отображая его информацию.
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что используется метод GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Извлекаем ID артиста из URL (например, из /artist/123)
	artistIDStr := r.URL.Path[len("/artist/"):]
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		InternalServerErrorHandler(w, r, err) // Обрабатываем некорректный ID
		return
	}

	// 2. Получаем информацию об артисте из API
	artist, err := api.GetArtist(artistID)
	if err != nil {
		BadRequestHandler(w, r, err)
		return
	}

	// 3. Загружаем шаблон artist.html для отображения информации об артисте
	tmpl, err := template.ParseFiles("web/templates/artist.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Подготавливаем данные для шаблона
	data := struct {
		Artist models.Artist
	}{
		Artist: artist,
	}

	// Выполняем рендеринг шаблона с данными и отправляем результат клиенту
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
