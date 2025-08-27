package importRoutes

import (
	"context"
	"encoding/json"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/sources"
	"garg/types"
	"garg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ImportGitRepo(c *gin.Context) {
	var body types.ApiRepositoryImportRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeInvalidRequestBody,
			Message: types.ApiErrorMessageInvalidRequestBody,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if body.RepositoryUrl == "" {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeRepositoryUrlRequired,
			Message: types.ApiErrorMessageRepositoryUrlRequired,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if !strings.HasPrefix(body.RepositoryUrl, "http://") && !strings.HasPrefix(body.RepositoryUrl, "https://") {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeRepositoryUrlBadSchema,
			Message: types.ApiErrorMessageRepositoryUrlBadSchema,
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	repoSource := strings.Split(body.RepositoryUrl, "/")[2]
	repoOwner := strings.Split(body.RepositoryUrl, "/")[3]
	repoName := strings.Split(body.RepositoryUrl, "/")[4]

	var source sources.Git
	err := source.Import(types.Repo{
		Url:  body.RepositoryUrl,
	})

	if err != nil {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeRepositoryCloneFailed,
			Message: types.ApiErrorMessageRepositoryCloneFailed,
			Error:   err.Error(),
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	client, _ := dbClient.GetClient()
	createdRepo, _ := client.CreateRepo(context.Background(), db.CreateRepoParams{
		Owner:       repoOwner,
		Name:        utils.RemoveDotGitExt(repoName),
		OriginalUrl: utils.AppendDotGitExt(body.RepositoryUrl),
		CreatedAt:   time.Now(),
		Source:      repoSource,
	})

	response := types.ApiRepositoryResponse{
		Id:          createdRepo.ID,
		Owner:       createdRepo.Owner,
		Name:        createdRepo.Name,
		CreatedAt:   createdRepo.CreatedAt.String(),
		OriginalUrl: createdRepo.OriginalUrl,
		Source:      createdRepo.Source,
	}

	c.JSON(http.StatusOK, response)
}
