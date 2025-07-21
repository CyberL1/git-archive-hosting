package reposRoutes

import (
	"context"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSingleRepo(c *gin.Context) {
	repoOwner := c.Param("owner")
	repoName := c.Param("repo")

	client, _ := dbClient.GetClient()
	repo, _ := client.GetRepoByFullName(context.Background(), db.GetRepoByFullNameParams{
		Owner: repoOwner,
		Name:  repoName,
	})

	response := types.ApiRepositoryResponse{
		Id:          repo.ID,
		Owner:       repo.Owner,
		Name:        repo.Name,
		OriginalUrl: repo.OriginalUrl,
		CreatedAt:   repo.CreatedAt.String(),
	}

	c.JSON(http.StatusOK, response)
}
