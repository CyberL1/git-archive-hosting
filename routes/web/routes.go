package web

import (
	webMiddlewares "garg/middlewares/web"
	repoRoutes "garg/routes/web/repo"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	repo := r.Group("/:source/:owner/:repo")

	repo.Use(webMiddlewares.GitSmartHTTPMiddleware())
	repo.GET("/HEAD", repoRoutes.GetHead)
	repo.GET("/info/refs", repoRoutes.GetInfoRefs)
	repo.POST("/git-upload-pack", repoRoutes.PostGitUploadPack)
}
