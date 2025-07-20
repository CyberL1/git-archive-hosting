package web

import (
	importRoutes "garg/routes/web/import"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", dashboard)
	r.HandleFunc("GET /import/git", importRoutes.ImportGitRepo)

	return r
}
