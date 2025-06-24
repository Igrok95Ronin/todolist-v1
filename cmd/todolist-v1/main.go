package main

import (
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/handlers"
	"github.com/Igrok95Ronin/todolist-v1.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.GetConfig()

	// Загружаем логгер
	logger := logging.GetLogger()

	// Создаем роутер
	router := httprouter.New()

	// Инициализируем обработчики (handlers) и передаем им зависимости
	handler := handlers.NewHandler(cfg, logger)
	handler.RegisterRoutes(router)

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
	log.Fatal(server.ListenAndServe())

}
