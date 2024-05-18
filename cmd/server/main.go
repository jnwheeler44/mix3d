package main

import (
	"mix3d/internal/routes"
)

func main() {

	// Register routes
	router := routes.RegisterRoutes()
	// Start the server
	router.Run(":8080") // Adjust port as needed
}
