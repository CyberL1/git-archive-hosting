package api

import (
	importRoutes "garg/routes/api/import"
	reposRoutes "garg/routes/api/repos"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/import/git", importRoutes.ImportGitRepo)
	r.GET("/repos", reposRoutes.GetAllRepos)
}
