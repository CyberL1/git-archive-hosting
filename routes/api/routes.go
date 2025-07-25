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

	repos := r.Group("/repos")
	repos.GET("/", reposRoutes.GetAllRepos)
	repos.GET("/:source", reposRoutes.GetReposBySource)
	repos.GET("/:source/:owner", reposRoutes.GetReposByOwner)
	repos.GET("/:source/:owner/:repo", reposRoutes.GetSingleRepo)
}
