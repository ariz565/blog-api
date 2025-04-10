package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post struct defines the data model for blog posts
type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// in-memory slice to store blog posts
var posts = []Post{
	{ID: "1", Title: "Hello World", Content: "This is my first blog post"},
}

// setupRoutes defines all endpoints and their handlers
func setupRoutes(r *gin.Engine) {
	// public routes
	r.POST("/login", loginHandler)

	// protected routes (requires JWT auth)
	protected := r.Group("/")
	protected.Use(authMiddleware()) // use middleware from auth.go
	{
		protected.GET("/posts", getPosts)
		protected.GET("/posts/:id", getPostByID)
		protected.POST("/posts", createPost)
		protected.PUT("/posts/:id", updatePost)
		protected.DELETE("/posts/:id", deletePost)
	}
}

// Handlers below 👇

// Get all blog posts
func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, posts)
}

// Get a single blog post by ID
func getPostByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range posts {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

// Create a new blog post
func createPost(c *gin.Context) {
	var newPost Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	posts = append(posts, newPost)
	c.JSON(http.StatusCreated, newPost)
}

// Update an existing blog post
func updatePost(c *gin.Context) {
	id := c.Param("id")
	var updatedPost Post
	if err := c.BindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	for i, p := range posts {
		if p.ID == id {
			posts[i] = updatedPost
			c.JSON(http.StatusOK, updatedPost)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

// Delete a blog post
func deletePost(c *gin.Context) {
	id := c.Param("id")
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
}
