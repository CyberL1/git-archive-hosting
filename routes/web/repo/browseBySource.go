package repoRoutes

import (
	"encoding/json"
	"fmt"
	"garg/types"
	"garg/utils"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BrowseBySource(c *gin.Context) {
	repoSource := c.Param("source")

	reposReq, _ := http.Get(fmt.Sprintf("%s/api/repos/%s", "http://localhost:8080", strings.ToLower(repoSource)))
	if reposReq.StatusCode == 404 {
		utils.RenderPage(c.Writer, "404", map[string]interface{}{
			"Title": "Page not found",
		})
		return
	}

	body, _ := io.ReadAll(reposReq.Body)

	var reposData []types.ApiRepositoryResponse
	json.Unmarshal(body, &reposData)

	utils.RenderPage(c.Writer, "repo/browseBySource", map[string]interface{}{
		"Title":        "Repositories imported from" + repoSource,
		"Source":       repoSource,
		"Repositories": reposData,
	})
}
