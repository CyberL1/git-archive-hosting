package main

import (
	"fmt"
	"garg/constants"
	dbClient "garg/db"
	"garg/routes/api"
	"garg/routes/web"
	"garg/utils"
	"net/http"
	"os"

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
