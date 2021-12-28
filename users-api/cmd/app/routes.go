package app

import (
	"github.com/ashtishad/bookstore/users-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func getRouteMappings(uh handlers.UserHandlers, r *gin.Engine) {
	r.GET("/users/:id", uh.GetById)

	r.GET("/users", uh.GetAll)
	r.GET("/users/search", uh.Search)
	r.POST("/users", uh.Create)
	r.PUT("/users/:id", uh.Update)
	r.DELETE("/users/:id", uh.Delete)
}
