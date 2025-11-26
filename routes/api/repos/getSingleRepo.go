package reposRoutes

import (
	db "garg/db/generated"
	"garg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSingleRepo(c *gin.Context) {
	repo := c.MustGet("repository").(db.Repo)

	response := types.ApiRepositoryResponse{
		Id:          repo.ID,
		Owner:       repo.Owner,
		Name:        repo.Name,
		OriginalUrl: repo.OriginalUrl,
		CreatedAt:   repo.CreatedAt.String(),
		Source:      repo.Source,
	}

	c.JSON(http.StatusOK, response)
}
