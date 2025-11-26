package reposRoutes

import (
	"garg/constants"
	"garg/types"
	"garg/utils"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/filemode"
)

func GetRepoContents(c *gin.Context) {
	repoSource := strings.ToLower(c.Param("source"))
	repoOwner := strings.ToLower(c.Param("owner"))
	repoName := utils.AppendDotGitExt(strings.ToLower(c.Param("repo")))

	openRepo, err := git.PlainOpen(filepath.Join(constants.RepositoriesDir, repoSource, repoOwner, repoName))
	if err != nil {
		response := types.ApiErrorResponse{
			Message: "Failed to open repository",
			Error:   err.Error(),
		}

		c.JSON(500, response)
		return
	}

	ref, err := openRepo.Head()
	if err != nil {
		response := types.ApiErrorResponse{
			Message: "Failed to get repository HEAD",
			Error:   err.Error(),
		}

		c.JSON(500, response)
		return
	}

	commit, err := openRepo.CommitObject(ref.Hash())
	if err != nil {
		response := types.ApiErrorResponse{
			Message: "Failed to get latest commit",
			Error:   err.Error(),
		}

		c.JSON(500, response)
		return
	}

	tree, err := commit.Tree()
	if err != nil {
		response := types.ApiErrorResponse{
			Message: "Failed to get repository tree",
			Error:   err.Error(),
		}

		c.JSON(500, response)
		return
	}

	path := strings.Trim(c.Param("path"), "/")

	if len(path) > 0 {
		tree, err = tree.Tree(path)
		if err != nil {
			response := types.ApiErrorResponse{
				Code:    types.ApiErrorCodeNotFound,
				Message: "Failed to get repository tree at specified path",
				Error:   err.Error(),
			}

			c.JSON(500, response)
			return
		}
	}

	var contents []types.ApiRepositoryContentsItemResponse
	for _, entry := range tree.Entries {
		var fileType string
		switch entry.Mode {
		case filemode.Dir:
			fileType = "dir"
		case filemode.Regular:
			fileType = "file"
		case filemode.Symlink:
			fileType = "symlink"
		default:
			fileType = "unknown"
		}

		var size int64 = 0
		if entry.Mode.IsFile() {
			blob, _ := openRepo.BlobObject(entry.Hash)
			size = blob.Size
		}

		item := types.ApiRepositoryContentsItemResponse{
			Name: entry.Name,
			Type: fileType,
			Size: size,
		}

		contents = append(contents, item)
	}

	c.JSON(200, contents)
}
