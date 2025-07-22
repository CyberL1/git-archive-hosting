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

func View(c *gin.Context) {
	repoOwner := c.Param("owner")
	repoName := c.Param("repo")

	if strings.HasSuffix(repoName, ".git") {
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s/%s", repoOwner, utils.RemoveDotGitExt(repoName)))
	}

	reposReq, _ := http.Get(fmt.Sprintf("%s/api/repos/%s/%s", "http://localhost:8080", repoOwner, repoName))
	if reposReq.StatusCode == 404 {
		utils.RenderPage(c.Writer, "404", map[string]interface{}{
			"Title": "Page not found",
		})
		return
	}

	body, _ := io.ReadAll(reposReq.Body)

	var repoData types.ApiRepositoryResponse
	json.Unmarshal(body, &repoData)

	utils.RenderPage(c.Writer, "repo/view", map[string]interface{}{
		"Title":      repoData.Owner + "/" + repoData.Name,
		"Repository": repoData,
	})
}
