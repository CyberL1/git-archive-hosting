package sources

import (
	"fmt"
	"garg/constants"
	"garg/types"
	"garg/utils"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type Git struct {
	Username string
	Password string
}

func (g *Git) Import(repo types.Repo) error {
	repoSource := strings.Split(repo.Url, "/")[2]
	repoOwner := strings.Split(repo.Url, "/")[3]
	repoName := strings.Split(repo.Url, "/")[4]

	fmt.Println("Importing repository:", repo.Url)

	cloneOptions := &git.CloneOptions{URL: repo.Url}

	if g.Username != "" && g.Password != "" {
		cloneOptions.Auth = &http.BasicAuth{
			Username: g.Username,
			Password: g.Password,
		}
	}

	_, err := git.PlainClone(filepath.Join(constants.RepositoriesDir, strings.ToLower(repoSource), strings.ToLower(repoOwner), utils.AppendDotGitExt(strings.ToLower(repoName))), true, cloneOptions)
	if err != nil {
		fmt.Println("Import failed:", err)
		return err
	}
	return nil
}
