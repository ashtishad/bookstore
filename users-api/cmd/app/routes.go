package app

import (
	"github.com/ashtishad/bookstore/users-api/internal/handlers"
)

func getRouteMappings() {
	r.GET("/users", handlers.GetAll)
	r.GET("/users/:id", handlers.GetById)
	r.GET("/users/search", handlers.Search)
	r.POST("/users", handlers.Create)
	r.PUT("/users/:id", handlers.Update)
	r.DELETE("/users/:id", handlers.Delete)
}
