package sources

import (
	"fmt"
	"garg/constants"
	"garg/types"
	"garg/utils"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

type Git struct{}

func (g *Git) Import(repo types.Repo) error {
	repoSource := strings.Split(repo.Url, "/")[2]
	repoOwner := strings.Split(repo.Url, "/")[3]
	repoName := strings.Split(repo.Url, "/")[4]

	fmt.Println("Importing repository:", repo.Url)

	_, err := git.PlainClone(filepath.Join(constants.RepositoriesDir, strings.ToLower(repoSource), strings.ToLower(repoOwner), utils.AppendDotGitExt(strings.ToLower(repoName))), true, &git.CloneOptions{
		URL: repo.Url,
	})
	if err != nil {
		fmt.Println("Import failed:", err)
		return err
	}
	return nil
}
