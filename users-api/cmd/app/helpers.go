package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func getRouteMappings() {
	//r.GET("/users", getUsers)
	//r.GET("/users/:id", getUser)
	//r.POST("/users", createUser)
	//r.PUT("/users/:id", updateUser)
	//r.DELETE("/users/:id", deleteUser)
}

// wait for interrupt signal to gracefully shut down the server with a timeout of 30 seconds.
func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// graceful shutdown
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	log.Println("Server gracefully stopped")
}
