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

	"gopkg.in/src-d/go-git.v4"
)

func ImportGitRepo(w http.ResponseWriter, r *http.Request) {
	var body types.ApiRepositoryImportRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeInvalidRequestBody,
			Message: types.ApiErrorMessageInvalidRequestBody,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if body.RepositoryUrl == "" {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeRepositoryUrlRequired,
			Message: types.ApiErrorMessageRepositoryUrlRequired,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	if !strings.HasPrefix(body.RepositoryUrl, "http://") && !strings.HasPrefix(body.RepositoryUrl, "https://") {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeRepositoryUrlBadSchema,
			Message: types.ApiErrorMessageRepositoryUrlBadSchema,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	repoOwner := strings.Split(body.RepositoryUrl, "/")[3]
	repoName := strings.Split(body.RepositoryUrl, "/")[4]

	_, err := git.PlainClone(filepath.Join(constants.RepositoriesDir, repoOwner, utils.AppendDotGitExt(repoName)), true, &git.CloneOptions{
		URL: body.RepositoryUrl,
	})

	if err != nil {
		response := types.ApiErrorResponse{
			Code:    types.ApiErrorCodeRepositoryCloneFailed,
			Message: types.ApiErrorMessageRepositoryCloneFailed,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	client, _ := dbClient.GetClient()
	createdRepo, _ := client.CreateRepo(context.Background(), db.CreateRepoParams{
		Owner:       repoOwner,
		Name:        utils.RemoveDotGitExt(repoName),
		OriginalUrl: utils.AppendDotGitExt(body.RepositoryUrl),
		CreatedAt:   time.Now(),
	})

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(createdRepo)
}
