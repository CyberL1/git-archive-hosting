package repoRoutes

import (
	"fmt"
	"garg/constants"
	"garg/utils"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostGitUploadPack(c *gin.Context) {
	repoOwner := c.Param("owner")
	repoName := c.Param("repo")

	repoPath := filepath.Join(constants.RepositoriesDir, strings.ToLower(repoOwner), utils.AppendDotGitExt(strings.ToLower(repoName)))

	c.Header("Content-Type", "application/x-git-upload-pack-result")
	c.Header("Cache-Control", "no-cache")

	cmd := exec.Command("git", "upload-pack", "--stateless-rpc", repoPath)
	cmd.Stdin = c.Request.Body
	cmd.Stdout = c.Writer
	cmd.Stderr = c.Writer

	err := cmd.Run()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("upload-pack failed: %v", err))
	}
}
