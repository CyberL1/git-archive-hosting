package importRoutes

import (
	"context"
	"encoding/json"
	"garg/constants"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/types"
	"garg/utils"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/src-d/go-git.v4"
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

	_, err := git.PlainClone(filepath.Join(constants.RepositoriesDir, strings.ToLower(repoSource), strings.ToLower(repoOwner), utils.AppendDotGitExt(strings.ToLower(repoName))), true, &git.CloneOptions{
		URL: body.RepositoryUrl,
	})

	if err != nil {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeRepositoryCloneFailed,
			Message: types.ApiErrorMessageRepositoryCloneFailed,
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

	c.JSON(http.StatusOK, createdRepo)
}
