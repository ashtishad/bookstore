package handlers

import (
	"github.com/ashtishad/bookstore/users-api/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlers struct {
	Service services.UserService
}

// GetById returns a user by id
func (u UserHandlers) GetById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "get by id",
	})
}

// GetAll returns all users, paginated
func (u UserHandlers) GetAll(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "get all",
	})
}

// Search users by name, paginated
func (u UserHandlers) Search(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "search",
	})
}

// Create a new user
func (u UserHandlers) Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "create",
	})
}

// Update an existing user
func (u UserHandlers) Update(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "update",
	})
}

// Delete an existing user
func (u UserHandlers) Delete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "delete",
	})
}
