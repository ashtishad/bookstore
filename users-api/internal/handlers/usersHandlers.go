package handlers

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/ashtishad/bookstore/users-api/internal/dto"
	"github.com/ashtishad/bookstore/users-api/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlers struct {
	Service services.UserService
}

// GetById returns a user by id
func (u UserHandlers) GetById(c *gin.Context) {
	id, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.AsStatus(), err)
		return
	}
	user, err := u.Service.GetById(id)
	if err != nil {
		c.JSON(err.AsStatus(), err)
		return
	}

	c.JSON(http.StatusOK, user)
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
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := lib.NewBadRequestError("invalid json body")
		c.JSON(restErr.AsStatus(), restErr)
		return
	}

	result, saveErr := u.Service.Create(req)
	if saveErr != nil {
		c.JSON(saveErr.AsStatus(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
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
