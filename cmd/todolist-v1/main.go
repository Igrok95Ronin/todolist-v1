package main

import (
	"fmt"
	"github.com/Igrok95Ronin/todolist-v1.git/internal/config"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.GetConfig()

	router := httprouter.New()
	router.GET("/", Home)

	start(router, cfg)
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Home"))
}

func start(router *httprouter.Router, cfg *config.Config) {
	const timeout = 15 * time.Second

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
	}

	fmt.Printf("Сервер запущен на %v", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
