package main

import (
	"context"
	"fmt"
	"garg/constants"
	dbClient "garg/db"
	db "garg/db/generated"
	"garg/routes/api"
	"garg/routes/web"
	"garg/utils"
	"net/http"
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

	if !utils.IsDevMode() {
		gin.SetMode("release")
	}

	r := gin.Default()

	web.NewHandler().RegisterRoutes(r.Group("/"))
	api.NewHandler().RegisterRoutes(r.Group("/api"))

	r.NoRoute(func(c *gin.Context) {
		utils.RenderPage(c.Writer, "404", map[string]interface{}{
			"Title": "Page not found",
		})
	})

	println("Git Archive Hosting started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
