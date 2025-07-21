package web

import (
	"encoding/json"
	"garg/types"
	"garg/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func dashboard(c *gin.Context) {
	reposReq, _ := http.Get("http://localhost:8080" + "/api/repos")
	body, _ := io.ReadAll(reposReq.Body)

	var reposResp []types.ApiRepositoryResponse
	json.Unmarshal(body, &reposResp)

	utils.RenderPage(c.Writer, "index", map[string]interface{}{
		"Title":        "Home",
		"Repositories": reposResp,
	})
}
