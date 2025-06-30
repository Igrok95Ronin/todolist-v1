package handlers

import (
	"encoding/json"
	"errors"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/httperror"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/models"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/service"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/httperrorsend"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/successful"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var emailErrors = []error{
	httperror.ErrEmailMissingAt,
	httperror.ErrEmailInvalidAtPos,
	httperror.ErrEmailTooShort,
	httperror.ErrEmailRegexCheckFail,
	httperror.ErrEmailRegexMismatch,
}

// UserHandler обрабатывает запросы, связанные с users
type UserHandler struct {
	service service.UserService
	logger  *logging.Logger
}

func NewUserHandler(service service.UserService, logger *logging.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var users models.Users

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		httperrorsend.WriteJSONError(w, httperror.ErrJSONDecode.Error(), http.StatusBadRequest)
		h.logger.Errorf("%s: %s", httperror.ErrJSONDecode, err)
		return
	}

	// UserExists проверяем есть ли пользователь и регистрирует нового пользователя
	if err := h.service.UserExists(ctx, users); err != nil {
		if errors.Is(err, httperror.ErrMissingFields) {
			httperrorsend.WriteJSONError(w, httperror.ErrMissingFields.Error(), http.StatusBadRequest)
			h.logger.Errorf("h <- %s: %s", httperror.ErrRegistrationDenied, err)
			return
		}

		// Проходимся по ошибкам
		for _, e := range emailErrors {
			if errors.Is(err, e) {
				httperrorsend.WriteJSONError(w, httperror.ErrInvalidEmailFormat.Error(), http.StatusBadRequest)
				h.logger.Errorf("h <- %s: %s", httperror.ErrInvalidEmailFormat, err)
				return
			}
		}

		h.logger.Errorf("h <- %s: %s", httperror.ErrInternalServer, err)
		httperrorsend.WriteJSONError(w, httperror.ErrUserExists.Error(), http.StatusConflict)
		return
	}

	// Отправляем ответ
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte(successful.UserRegisteredSuccess))
	if err != nil {
		h.logger.Errorf("%s: %s", httperror.ErrResponseEncoding, err)
	}
	h.logger.Infof("%s: %s", users.UserName, successful.UserRegisteredSuccess)
}
