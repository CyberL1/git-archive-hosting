package reposRoutes

import (
	"context"
	"encoding/json"
	dbClient "garg/db"
	"garg/types"
	"net/http"
)

func GetAllRepos(w http.ResponseWriter, r *http.Request) {
	client, _ := dbClient.GetClient()
	repos, _ := client.ListRepos(context.Background())

	var response []types.ApiRepositoryResponse
	for _, repo := range repos {
		repo := types.ApiRepositoryResponse{
			Id:          repo.ID,
			Owner:       repo.Owner,
			Name:        repo.Name,
			OriginalUrl: repo.OriginalUrl,
			CreatedAt: repo.CreatedAt.String(),
		}

		response = append(response, repo)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
