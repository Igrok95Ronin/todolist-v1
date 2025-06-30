package httperrorsend

import (
	"encoding/json"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"

	"net/http"
)

// Структура для возврата ошибок
type ErrorResponse struct {
	Code     int    `json:"code"`
	CodeText string `json:"codeText"`
	Message  string `json:"message"`
}

// Функция для возврата ошибок в формате JSON
func WriteJSONError(w http.ResponseWriter, message string, code int) {
	logger := logging.GetLogger()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Code:     code,
		CodeText: http.StatusText(code),
		Message:  message,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		logger.Error("Ошибка при кодировании JSON-ответа: ", err)
	}
}
