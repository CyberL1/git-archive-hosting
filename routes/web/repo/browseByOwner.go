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

func BrowseByOwner(c *gin.Context) {
	repoSource := c.Param("source")
	repoOwner := c.Param("owner")

	reposReq, _ := http.Get(fmt.Sprintf("%s/api/repos/%s/%s", "http://localhost:8080", strings.ToLower(repoSource), strings.ToLower(repoOwner)))
	if reposReq.StatusCode == 404 {
		utils.RenderPage(c.Writer, "404", map[string]interface{}{
			"Title": "Page not found",
		})
		return
	}

	body, _ := io.ReadAll(reposReq.Body)

	var reposData []types.ApiRepositoryResponse
	json.Unmarshal(body, &reposData)

	utils.RenderPage(c.Writer, "repo/browseByOwner", map[string]interface{}{
		"Title":        "Repositories imported from" + repoSource + "owned by" + repoOwner,
		"Source":       repoSource,
		"Owner":        repoOwner,
		"Repositories": reposData,
	})
}
