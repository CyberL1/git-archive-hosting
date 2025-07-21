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
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func GetInfoRefs(c *gin.Context) {
	if c.Query("service") != "git-upload-pack" {
		c.String(http.StatusBadRequest, "Missing or invalid service")
		return
	}

	repoOwner := c.Param("owner")
	repoName := c.Param("repo")

	repoPath := filepath.Join(constants.RepositoriesDir, strings.ToLower(repoOwner), utils.AppendDotGitExt(strings.ToLower(repoName)))

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open repository: %v", err))
		return
	}

	refs, err := repo.References()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to get references: %v", err))
		return
	}

	var refsList []string
	err = refs.ForEach(func(ref *plumbing.Reference) error {
		refsList = append(refsList, fmt.Sprintf("%s %s", ref.Hash().String(), ref.Name().String()))
		return nil
	})

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to iterate references: %v", err))
		return
	}

	c.Header("Content-Type", "application/x-git-upload-pack-advertisement")
	c.Header("Cache-Control", "no-cache")

	serviceLine := "# service=git-upload-pack\n"
	length := len(serviceLine) + 4
	serviceHeader := fmt.Sprintf("%04x%s0000", length, serviceLine)

	cmd := exec.Command("git", "upload-pack", "--stateless-rpc", "--advertise-refs", repoPath)
	output, err := cmd.Output()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to execute git-upload-pack: %v", err))
		return
	}

	c.Writer.Write([]byte(serviceHeader))
	c.Writer.Write(output)
}
