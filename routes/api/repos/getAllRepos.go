package reposRoutes

import (
	"context"
	dbClient "garg/db"
	"garg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRepos(c *gin.Context) {
	client, _ := dbClient.GetClient()
	repos, _ := client.ListRepos(context.Background())

	var response []types.ApiRepositoryResponse
	for _, repo := range repos {
		repo := types.ApiRepositoryResponse{
			Id:          repo.ID,
			Owner:       repo.Owner,
			Name:        repo.Name,
			OriginalUrl: repo.OriginalUrl,
			CreatedAt:   repo.CreatedAt.String(),
		}

		response = append(response, repo)
	}

	c.JSON(http.StatusOK, response)
}
