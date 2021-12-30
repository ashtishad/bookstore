package app

import (
	"github.com/ashtishad/bookstore/users-api/internal/rest"
	"github.com/gin-gonic/gin"
)

func getRouteMappings(uh rest.UserHandlers, r *gin.Engine) {
	r.GET("/users/:id", uh.GetById)
	r.POST("/users", uh.Create)
	r.PUT("/users/:id", uh.Update)

	r.DELETE("/users/:id", uh.Delete)
	r.GET("/users", uh.GetAll)
	r.GET("/users/search", uh.Search)
}
