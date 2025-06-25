package main

import (
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/handlers"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/repository"
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

	// Инициализируем обработчики (handlers) и передаем им зависимости
	handler, err := handlers.NewHandler(
		handlers.WithConfig(cfg),
		handlers.WithLogger(logger),
	)
	if err != nil {
		logger.Error(err)
	}
	handler.RegisterRoutes(router)

	// Запускаем сервер
	start(router, cfg, logger)
}

func start(router *httprouter.Router, cfg *config.Config, logger *logging.Logger) {
	const timeout = 15 * time.Second

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
	}

	logger.Infof("Сервер запущен на %v", cfg.Port)
	logger.Fatal(server.ListenAndServe())

}
