package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var r = gin.Default()

func Start() {
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		IdleTimeout:    100 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	getRouteMappings()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			return
		}
	}()

	gracefulShutdown(srv)
}
