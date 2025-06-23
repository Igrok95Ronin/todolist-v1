package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	router := httprouter.New()
	router.GET("/", Home)

	start(router)
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Home"))
}

func start(router *httprouter.Router) {
	const timeout = 15 * time.Second

	server := &http.Server{
		Addr:         "8080",
		Handler:      router,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout,
	}

	fmt.Printf("Сервер запущен на 8080")
	log.Fatal(server.ListenAndServe())
}
