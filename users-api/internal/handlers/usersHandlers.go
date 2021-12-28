package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetById returns a user by id
func GetById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "get by id",
	})
}

// GetAll returns all users, paginated
func GetAll(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "get all",
	})
}

// Search users by name, paginated
func Search(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "search",
	})
}

// Create a new user
func Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "create",
	})
}

// Update an existing user
func Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "update",
	})
}

// Delete an existing user
func Delete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "delete",
	})
}
