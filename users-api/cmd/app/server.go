package app

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var r = gin.New()

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

	r.Use(gin.LoggerWithFormatter(lib.Logger))

	r.Use(gin.CustomRecovery(lib.Recover))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			return
		}
	}()

	lib.GracefulShutdown(srv)
}
