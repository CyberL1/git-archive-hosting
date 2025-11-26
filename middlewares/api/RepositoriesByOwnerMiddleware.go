package apiMiddlewares

import (
	"context"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RepositoriesByOwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		repoSource := c.Param("source")
		repoOwner := c.Param("owner")

		client, _ := dbClient.GetClient()
		repositories, _ := client.ListReposBySourceAndOwner(context.Background(), db.ListReposBySourceAndOwnerParams{
			Source: repoSource,
			Owner:  repoOwner,
		})

		if len(repositories) < 1 {
			response := types.ApiErrorResponse{
				Code:    types.ApiErrorCodeNotFound,
				Message: types.ApiErrorMessageOwnerNotFound,
			}

			c.JSON(http.StatusNotFound, response)
			c.Abort()
			return
		}

		c.Set("repositories", repositories)
		c.Next()
	}
}
