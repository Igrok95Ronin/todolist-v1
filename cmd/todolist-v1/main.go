package main

import (
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/handlers"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/middleware"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/repository"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/service"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func main() {
	// Подтягиваем конфигурацию
	cfg := config.GetConfig()

	// Подтягиваем логгер
	logger := logging.GetLogger()

	// Инициализируем базу данных (в слое repository)
	db := repository.NewDB(
		repository.WithConfig(cfg),
		repository.WithLogger(logger),
	)
	sqlDB, err := db.Connect()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Подтягиваем миграции
	if err = repository.InitSchema(sqlDB); err != nil {
		logger.Fatal(err)
	}

	// Создаем роутер
	router := httprouter.New()

	// Слои
	userRepo := repository.NewUserRepository(sqlDB)
	userSrv := service.NewUserService(userRepo, cfg)

	// Инициализируем обработчики (handlers) и передаем им зависимости
	handler, err := handlers.NewHandler(
		handlers.WithConfig(cfg),
		handlers.WithLogger(logger),
		handlers.WithDB(sqlDB),
		handlers.WithUserRepo(userRepo),
		handlers.WithUserSrv(userSrv),
	)
	if err != nil {
		logger.Error(err)
	}
	handler.RegisterRoutes(router)

	// Обработка cors, Context
	corsHandler := middleware.CorsSettings().Handler(middleware.RequestContext(router))

	// Запускаем сервер
	start(corsHandler, cfg, logger)

}

func start(router http.Handler, cfg *config.Config, logger *logging.Logger) {
	const timeout = 15 * time.Second

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
	}

	logger.Infof("Сервер запущен на порту: %v", cfg.Port)
	logger.Fatal(server.ListenAndServe())

}
