package reposRoutes

import (
	"context"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/types"
	"garg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSingleRepo(c *gin.Context) {
	repoOwner := c.Param("owner")
	repoName := c.Param("repo")

	client, _ := dbClient.GetClient()
	repo, err := client.GetRepoByFullName(context.Background(), db.GetRepoByFullNameParams{
		Owner: repoOwner,
		Name:  utils.RemoveDotGitExt(repoName),
	})

	if err != nil {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeNotFound,
			Message: types.ApiErrorMessageRepositoryNotFound,
		}

		c.JSON(http.StatusNotFound, response)
		return
	}

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
