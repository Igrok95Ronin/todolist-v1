package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("register"))
	fmt.Println("register")
}
