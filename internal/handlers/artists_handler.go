package handlers

import (
	"groupie-trackers/api"
	"groupie-trackers/internal/models"
	"net/http"
	"text/template"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/artists" {
		NotFoundHandler(w, r)
		return
	}
	// Проверяем, что используется метод GET
	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r)
		return
	}

	// Получаем данные об артистах из API
	artists, err := api.GetArtists()
	if err != nil {
		InternalServerErrorHandler(w, r, err) // Обрабатываем возможную ошибку
		return
	}

	// Загружаем шаблон index.html для отображения списка артистов
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	// Подготавливаем данные для шаблона
	data := struct {
		Artists []models.Artist
	}{
		Artists: artists,
	}

	// Выполняем рендеринг шаблона с данными и отправляем результат клиенту
	err = tmpl.Execute(w, data)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}
}
