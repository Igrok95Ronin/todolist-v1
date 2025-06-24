package handlers

import (
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Handler управляет роутами
type Handler struct {
	cfg    *config.Config
	logger *logging.Logger
}

// NOTE: добавить инкопсюляцию и функциональные параметры
// NewHandler создаёт новый обработчик
func NewHandler(cfg *config.Config, logger *logging.Logger) *Handler {
	h := &Handler{}
	h.SetConfig(cfg)
	h.SetLogger(logger)
	return h
}

// Set Сеттеры с логикой
func (h *Handler) SetConfig(cfg *config.Config) {
	h.cfg = cfg
}

func (h *Handler) SetLogger(logger *logging.Logger) {
	h.logger = logger
}

// Get Геттеры
func (h *Handler) Cfg() *config.Config {
	return h.cfg
}

func (h *Handler) Logger() *logging.Logger {
	return h.logger
}

// RegisterRoutes регистрирует маршруты
func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	router.GET("/", h.Home)
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Home"))

}
