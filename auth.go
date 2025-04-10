package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// User struct defines a mock user model for authentication
type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // In real apps, store hashed passwords!
}

// Dummy in-memory user list — this simulates user records from a database
var users = []User{
	{Username: "admin", Password: "password"}, // default user
}

// loginHandler handles login requests and returns a JWT if credentials are valid
func loginHandler(c *gin.Context) {
	var creds User
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Loop through dummy users and check for matching credentials
	for _, u := range users {
		if u.Username == creds.Username && u.Password == creds.Password {
			// If valid, generate JWT token
			token, err := generateJWT(u.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
				return
			}

			// Return the token in response
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}

	// If no match found
	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
}

// generateJWT creates and signs a JWT token using the globally loaded secret
func generateJWT(username string) (string, error) {
	// Define token claims (payload data)
	claims := jwt.MapClaims{
		"username": username,                              // include username
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // expire in 24 hours
	}

	// Create the token using HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret loaded from .env (defined in main.go)
	return token.SignedString(jwtSecret)
}

// authMiddleware is a Gin middleware that protects routes by validating JWTs
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the token from the Authorization header
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Make sure the token's signing method is what we expect
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		// Token is valid, continue with the request
		c.Next()
	}
}
