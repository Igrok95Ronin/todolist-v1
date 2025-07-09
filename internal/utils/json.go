package utils

import (
	"encoding/json"
	"net/http"
)

// JSONDecoder — обобщённая (generic) функция для декодирования JSON-запроса в любую структуру.
// T — это универсальный тип, например: dto.RegisterRequest, dto.LoginRequest и т.д.
func JSONDecoder[T any](r *http.Request, target *T) error {
	// Создаём JSON-декодер на основе тела запроса
	decoder := json.NewDecoder(r.Body)

	// Декодируем JSON в переданный указатель `target`.
	return decoder.Decode(target)
}
