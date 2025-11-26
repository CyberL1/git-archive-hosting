package apiMiddlewares

import (
	"context"
	dbClient "garg/db"
	"garg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RepositoriesBySourceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		source := c.Param("source")

		client, _ := dbClient.GetClient()
		repositories, _ := client.ListReposBySource(context.Background(), source)

		if len(repositories) < 1 {
			response := types.ApiErrorResponse{
				Code:    types.ApiErrorCodeNotFound,
				Message: types.ApiErrorMessageSourceNotFound,
			}

			c.JSON(http.StatusNotFound, response)
			c.Abort()
			return
		}

		c.Set("repositories", repositories)
		c.Next()
	}
}
