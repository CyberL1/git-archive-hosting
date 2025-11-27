package reposRoutes

import (
	"encoding/base64"
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

	if path == "" {
		path = "."
	}

	if path != "." {
		entry, err := tree.FindEntry(path)
		if err != nil {
			response := types.ApiErrorResponse{
				Code:    types.ApiErrorCodeNotFound,
				Message: "Failed to find entry at specified path",
				Error:   err.Error(),
			}

			c.JSON(500, response)
			return
		}

		if entry.Mode.IsFile() {
			file, _ := tree.File(path)
			fileContent, _ := file.Contents()

			response := types.ApiRepositoryContentsItemResponse{
				Name:    entry.Name,
				Type:    getFileType(entry.Mode),
				Size:    file.Size,
				Content: base64.RawStdEncoding.EncodeToString([]byte(fileContent)),
			}

			c.JSON(200, response)
			return
		}
	}

	if path != "." {
		tree, _ = tree.Tree(path)
	}

	var response []types.ApiRepositoryContentsItemResponse

	for _, entry := range tree.Entries {
		var size int64 = 0
		if entry.Mode.IsFile() {
			blob, _ := openRepo.BlobObject(entry.Hash)
			size = blob.Size
		}

		item := types.ApiRepositoryContentsItemResponse{
			Name: entry.Name,
			Type: getFileType(entry.Mode),
			Size: size,
		}

		response = append(response, item)
	}

	c.JSON(200, response)
}

func getFileType(mode filemode.FileMode) string {
	var fileType string
	switch mode {
	case filemode.Dir:
		fileType = "dir"
	case filemode.Regular:
		fileType = "file"
	case filemode.Symlink:
		fileType = "symlink"
	case filemode.Submodule:
		fileType = "submodule"
	default:
		fileType = "unknown"
	}

	return fileType
}
