package handlers

import (
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

// Handler управляет роутами
type Handler struct {
	cfg    *config.Config
	logger *logging.Logger
}

type HandlerOption func(*Handler)

// NOTE: добавить инкопсюляцию и функциональные параметры
// NewHandler создаёт новый обработчик
func NewHandler(option ...HandlerOption) (*Handler, error) {
	h := &Handler{}

	for _, opt := range option {
		opt(h)
	}

	//if h.logger == nil {
	//	return nil, fmt.Errorf("logger is required")
	//}
	//if h.cfg == nil {
	//	return nil, fmt.Errorf("config is required")
	//}

	return h, nil
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

// With Функции-опции
func WithConfig(cfg *config.Config) HandlerOption {
	return func(h *Handler) {
		h.SetConfig(cfg)
	}
}

func WithLogger(logger *logging.Logger) HandlerOption {
	return func(h *Handler) {
		h.SetLogger(logger)
	}
}

// RegisterRoutes регистрирует маршруты
func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	router.GET("/register", h.register)
}
