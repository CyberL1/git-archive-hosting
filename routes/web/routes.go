package web

import (
	importRoutes "garg/routes/web/import"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/", dashboard)
	r.GET("/import/git", importRoutes.ImportGitRepo)
}
