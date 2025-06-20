package routes

import (
	"fmt"
	"net/http"
)

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("POST /")
	r.HandleFunc("GET /")
}
