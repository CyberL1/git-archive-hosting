package api

import (
	importRoutes "garg/routes/api/import"
	reposRoutes "garg/routes/api/repos"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /import/git", importRoutes.ImportGitRepo)
	r.HandleFunc("GET /repos", reposRoutes.GetAllRepos)

	return r
}
