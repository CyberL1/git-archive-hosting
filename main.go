package main

import (
	"context"
	"fmt"
	"garg/constants"
	dbClient "garg/db"
	"garg/frontend"
	"garg/routes/api"
	"garg/routes/web"
	"garg/utils"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	// Create the data directory if it doesn't exist
	if _, err := os.Stat(constants.DataDir); os.IsNotExist(err) {
		if err := os.Mkdir(constants.DataDir, 0755); err != nil {
			fmt.Println("Error creating data directory:", err)
			os.Exit(1)
		}
	}

	// Migrate DB
	dbClient.Migrate()

	// Set non-mirror repos as mirrors
	fmt.Println("Converting non-mirror repositories to mirrors")
	client, _ := dbClient.GetClient()
	repos, _ := client.ListRepos(context.Background())

	for _, repo := range repos {
		repoPath := filepath.Join(constants.RepositoriesDir, repo.Source, strings.ToLower(repo.Owner), utils.AppendDotGitExt(strings.ToLower(repo.Name)))

		openRepo, err := git.PlainOpen(repoPath)
		if err != nil {
			fmt.Printf("Failed to open repo %s: %v\n", repoPath, err)
			continue
		}

		origin, _ := openRepo.Remote("origin")

		if origin.Config().Fetch[0] == config.RefSpec("+refs/*:refs/*") && origin.Config().Mirror {
			fmt.Printf("Skipping %s as it is already a mirror\n", repoPath)
			continue
		}

		fmt.Printf("Setting %s as a mirror\n", repoPath)

		branches, _ := openRepo.Branches()

		branches.ForEach(func(r *plumbing.Reference) error {
			openRepo.DeleteBranch(r.Name().Short())
			return nil
		})

		openRepo.DeleteRemote(origin.Config().Name)

		newRemote := &config.RemoteConfig{
			Name:   origin.Config().Name,
			URLs:   origin.Config().URLs,
			Fetch:  []config.RefSpec{config.RefSpec("+refs/*:refs/*")},
			Mirror: true,
		}

		openRepo.CreateRemote(newRemote)
	}

	if utils.IsDevMode() {
		os.Setenv("DOMAIN", "localhost:8080")
	} else {
		gin.SetMode("release")
	}

	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		if os.Getenv("DEV_MODE") == "true" {
			frontendUrl, _ := url.Parse("http://localhost:5173")
			httputil.NewSingleHostReverseProxy(frontendUrl).ServeHTTP(c.Writer, c.Request)
		} else {
			path := c.Request.URL.Path[1:]

			if _, err := frontend.BuildDir.ReadFile(filepath.Join("build", path)); err != nil {
				path = "index.html"
			}

			fileContent, _ := frontend.BuildDir.ReadFile(filepath.Join("build", path))

			c.Header("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
			c.String(200, string(fileContent))
		}
	})

	web.NewHandler().RegisterRoutes(r.Group("/"))
	api.NewHandler().RegisterRoutes(r.Group("/api"))

	println("Git Archive Hosting started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
