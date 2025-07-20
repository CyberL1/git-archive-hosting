package web

import (
	"encoding/json"
	"garg/types"
	"garg/utils"
	"io"
	"net/http"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	reposReq, _ := http.Get("http://localhost:8080" + "/api/repos")
	body, _ := io.ReadAll(reposReq.Body)

	var reposResp []types.ApiRepositoryResponse
	json.Unmarshal(body, &reposResp)

	utils.RenderPage(w, "index", map[string]interface{}{
		"Title":        "Home",
		"Repositories": reposResp,
	})
}
