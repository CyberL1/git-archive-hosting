package api

import (
	apiMiddlewares "garg/middlewares/api"
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

	repos.Use(apiMiddlewares.RepositoriesBySourceMiddleware())
	repos.GET("/:source", reposRoutes.GetReposBySource)

	repos.Use(apiMiddlewares.RepositoriesByOwnerMiddleware())
	repos.GET("/:source/:owner", reposRoutes.GetReposByOwner)

	repo := repos.Group("/:source/:owner/:repo")

	repo.Use(apiMiddlewares.SingleRepositoryMiddleware())
	repo.GET("/", reposRoutes.GetSingleRepo)

	repo.GET("/contents", reposRoutes.GetRepoContents)
	repo.GET("/contents/*path", reposRoutes.GetRepoContents)
}
