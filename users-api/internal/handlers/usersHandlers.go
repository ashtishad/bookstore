package handlers

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/ashtishad/bookstore/users-api/internal/dto"
	"github.com/ashtishad/bookstore/users-api/pkg/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandlers struct {
	Service services.UserService
}

// GetById returns a user by id
func (u UserHandlers) GetById(c *gin.Context) {
	id, err := getUserId(c.Param("id"))
	if err != nil {
		log.Printf("Error while getting user id : %v", err)
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

// Create a new user
func (u UserHandlers) Create(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error while binding request: %v", err.Error())
		restErr := lib.NewBadRequestError("invalid json body")
		c.JSON(restErr.AsStatus(), restErr)
		return
	}

	result, saveErr := u.Service.Create(req)
	if saveErr != nil {
		log.Printf("Error while creating user: %v", saveErr)
		c.JSON(saveErr.AsStatus(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Update an existing user
func (u UserHandlers) Update(c *gin.Context) {
	id, err := getUserId(c.Param("id"))
	if err != nil {
		log.Printf("Error while getting user id: %v", err)
		c.JSON(err.AsStatus(), err)
		return
	}

	var req dto.UserUpdateRequest
	req.SetId(id)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error while binding request: %v", err.Error())
		restErr := lib.NewBadRequestError("invalid json body")
		c.JSON(restErr.AsStatus(), restErr)
		return
	}

	result, saveErr := u.Service.Update(req)

	if saveErr != nil {
		log.Printf("Error while updating user: %v", saveErr)
		c.JSON(saveErr.AsStatus(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// TODO: Delete an existing user
func (u UserHandlers) Delete(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "delete",
	})
}

// TODO: GetAll returns all users, paginated
func (u UserHandlers) GetAll(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "get all",
	})
}

// TODO: Search users by name, paginated
func (u UserHandlers) Search(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "search",
	})
}
