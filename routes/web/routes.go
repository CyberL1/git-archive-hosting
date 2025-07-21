package web

import (
	webMiddlewares "garg/middlewares/web"
	importRoutes "garg/routes/web/import"
	repoRoutes "garg/routes/web/repo"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/", dashboard)
	r.GET("/import/git", importRoutes.ImportGitRepo)

	repo := r.Group("/:owner/:repo")
	repo.GET("/", repoRoutes.View)

	repo.Use(webMiddlewares.GitSmartHTTPMiddleware())
	repo.GET("/HEAD", repoRoutes.GetHead)
	repo.GET("/info/refs", repoRoutes.GetInfoRefs)
	repo.POST("/git-upload-pack", repoRoutes.PostGitUploadPack)
}
