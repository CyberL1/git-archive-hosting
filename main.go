package main

import (
	"fmt"
	"garg/constants"
	dbClient "garg/db"
	"garg/routes/api"
	"garg/routes/web"
	"net/http"
	"os"
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

	r := http.NewServeMux()

	r.Handle("/", web.NewHandler().RegisterRoutes())
	r.Handle("/api/", http.StripPrefix("/api", api.NewHandler().RegisterRoutes()))

	println("Git Archive Hosting started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
