package handlers

import (
	"database/sql"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/repository"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/service"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

// Handler управляет роутами
type Handler struct {
	cfg      *config.Config
	logger   *logging.Logger
	db       *sql.DB
	userRepo repository.UserRepository
	userSvc  service.UserService
}

type HandlerOption func(*Handler)

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

// NOTE: Set Сеттеры с логикой
func (h *Handler) SetConfig(cfg *config.Config) {
	h.cfg = cfg
}

func (h *Handler) SetLogger(logger *logging.Logger) {
	h.logger = logger
}

func (h *Handler) SetDB(db *sql.DB) {
	h.db = db
}

func (h *Handler) SetUserRepo(userRepo repository.UserRepository) {
	h.userRepo = userRepo
}

func (h *Handler) SetUserSrc(userSvc service.UserService) {
	h.userSvc = userSvc
}

// NOTE: Get Геттеры
func (h *Handler) Cfg() *config.Config {
	return h.cfg
}

func (h *Handler) Logger() *logging.Logger {
	return h.logger
}

func (h *Handler) DB() *sql.DB {
	return h.db
}

func (h *Handler) UserRepo() repository.UserRepository {
	return h.userRepo
}

func (h *Handler) UserSrc() service.UserService {
	return h.userSvc
}

// NOTE: With Функции-опции
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

func WithDB(db *sql.DB) HandlerOption {
	return func(h *Handler) {
		h.SetDB(db)
	}
}

func WithUserRepo(userRepo repository.UserRepository) HandlerOption {
	return func(h *Handler) {
		h.SetUserRepo(userRepo)
	}
}

func WithUserSrv(userSrv service.UserService) HandlerOption {
	return func(h *Handler) {
		h.SetUserSrc(userSrv)
	}
}

// RegisterRoutes регистрирует маршруты
func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	userHandler := NewUserHandler(h.userSvc, h.logger)

	router.POST("/register", userHandler.register)

}
