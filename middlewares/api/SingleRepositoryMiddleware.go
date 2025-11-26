package apiMiddlewares

import (
	"context"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/types"
	"garg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SingleRepositoryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		repoSource := c.Param("source")
		repoOwner := c.Param("owner")
		repoName := c.Param("repo")

		client, _ := dbClient.GetClient()
		repository, err := client.GetRepoByFullName(context.Background(), db.GetRepoByFullNameParams{
			Source: repoSource,
			Owner:  repoOwner,
			Name:   utils.RemoveDotGitExt(repoName),
		})

		if err != nil {
			response := types.ApiErrorResponse{
				Code:    types.ApiErrorCodeNotFound,
				Message: types.ApiErrorMessageRepositoryNotFound,
			}

			c.JSON(http.StatusNotFound, response)
			c.Abort()
			return
		}

		c.Set("repository", repository)
		c.Next()
	}
}
