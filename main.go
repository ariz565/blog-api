package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var jwtSecret []byte // globally accessible variable

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get JWT_SECRET and set it globally
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not set in .env file")
	}
	jwtSecret = []byte(secret)

	r := gin.Default()
	setupRoutes(r)
	r.Run(":8080")
}
