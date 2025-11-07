package main

import (
	"context"
	"fmt"
	"garg/constants"
	dbClient "garg/db"
	db "garg/db/generated"
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

	// Migrate repos to the new system (owner/repo -> source/owner/repo)
	fmt.Println("Migrating repos to the new archive system")
	client, _ := dbClient.GetClient()
	repos, _ := client.ListRepos(context.Background())

	for _, repo := range repos {
		repoSource := strings.Split(repo.OriginalUrl, "/")[2]

		oldOwnerPath := filepath.Join(constants.RepositoriesDir, strings.ToLower(repo.Owner))
		newOwnerPath := filepath.Join(constants.RepositoriesDir, repoSource, strings.ToLower(repo.Owner))

		if err := os.MkdirAll(filepath.Join(constants.RepositoriesDir, repoSource), 0755); err != nil {
			fmt.Printf("Failed to migrate owner %s: %v\n", repo.Owner, err)
			continue
		}

		if err := os.Rename(oldOwnerPath, newOwnerPath); err != nil {
			fmt.Printf("Failed to migrate owner %s: %v\n", repo.Owner, err)
			continue
		}

		client.SetRepoSource(context.Background(), db.SetRepoSourceParams{
			Source: repoSource,
			Owner:  repo.Owner,
			Name:   repo.Name,
		})
		fmt.Printf("Migrated %s\n", repo.Owner)
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
