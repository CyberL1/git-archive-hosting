package reposRoutes

import (
	"context"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReposByOwner(c *gin.Context) {
	client, _ := dbClient.GetClient()
	repos, _ := client.ListReposBySourceAndOwner(context.Background(), db.ListReposBySourceAndOwnerParams{
		Source: c.Param("source"),
		Owner: c.Param("owner"),
	})

	var response []types.ApiRepositoryResponse
	for _, repo := range repos {
		repo := types.ApiRepositoryResponse{
			Id:          repo.ID,
			Owner:       repo.Owner,
			Name:        repo.Name,
			OriginalUrl: repo.OriginalUrl,
			CreatedAt:   repo.CreatedAt.String(),
			Source:      repo.Source,
		}

		response = append(response, repo)
	}

	c.JSON(http.StatusOK, response)
}
