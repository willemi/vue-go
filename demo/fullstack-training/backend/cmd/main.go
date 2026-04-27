package main

import (
	"fullstack-backend/config"
	"fullstack-backend/server"
)

func main() {
	// Initialize configuration
	config.Init()

	// Initialize server
	r := server.New()

	// Start server
	r.Run(":8080")
}