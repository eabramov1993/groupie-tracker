package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-trackers/internal/models" // Путь к вашим моделям данных
)

// getRelations получает информацию о связях артиста по его ID
func GetRelations(artistID int) (models.Relation, error) {
	// Формируем URL для запроса к API, подставляя artistID
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", artistID)

	// Выполняем HTTP GET запрос по сформированному URL
	resp, err := http.Get(url)
	if err != nil {
		// В случае ошибки возвращаем пустую структуру Relation и ошибку
		return models.Relation{}, fmt.Errorf("failed to get relations: %v", err)
	}
	// Закрываем тело ответа после завершения функции
	defer resp.Body.Close()

	// Создаем переменную для хранения данных о связях
	var relation models.Relation

	// Декодируем JSON ответ в переменную relation
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		// В случае ошибки декодирования возвращаем пустую структуру Relation и ошибку
		return models.Relation{}, fmt.Errorf("failed to decode relations: %v", err)
	}

	// Возвращаем заполненную структуру Relation и nil (отсутствие ошибки)
	return relation, nil
}
