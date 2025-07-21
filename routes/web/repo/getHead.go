package repoRoutes

import (
	"fmt"
	"garg/constants"
	"garg/utils"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/src-d/go-git.v4"
)

func GetHead(c *gin.Context) {
	repoOwner := c.Param("owner")
	repoName := c.Param("repo")

	// Only allow smart HTTP requests:
	if !strings.Contains(c.Request.UserAgent(), "git/") {
		c.String(http.StatusForbidden, "Access denied: Not a git client")
		return
	}

	// Accept only smart HTTP query param or content type
	if c.Query("service") != "git-upload-pack" && c.ContentType() != "application/x-git-upload-pack-request" {
		c.String(http.StatusForbidden, "Access denied: Not a smart HTTP request")
		return
	}

	repo, err := git.PlainOpen(filepath.Join(constants.RepositoriesDir, repoOwner, utils.AppendDotGitExt(repoName)))
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open repository: %v", err))
		return
	}

	head, err := repo.Head()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to get HEAD: %v", err))
		return
	}

	c.String(http.StatusOK, head.Name().String())
}
