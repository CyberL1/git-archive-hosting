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

	repo, err := git.PlainOpen(filepath.Join(constants.RepositoriesDir, strings.ToLower(repoOwner), utils.AppendDotGitExt(strings.ToLower(repoName))))
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
