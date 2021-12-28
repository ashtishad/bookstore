package app

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/ashtishad/bookstore/users-api/data"
	"github.com/ashtishad/bookstore/users-api/internal/handlers"
	"github.com/ashtishad/bookstore/users-api/pkg/domain"
	"github.com/ashtishad/bookstore/users-api/pkg/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
	var r = gin.New()

	// db connection pool config
	db := data.GetDbClient()
	defer func() {
		_ = db.Close()
		log.Println("DB connection pool closed")
	}()
	userDbConn := domain.NewUserRepoDb(db)

	// wire up handlers
	uh := handlers.UserHandlers{Service: services.NewUserService(userDbConn)}

	// Server Config
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		IdleTimeout:    100 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// route url mappings
	getRouteMappings(uh, r)

	// custom logger middleware
	r.Use(gin.LoggerWithFormatter(lib.Logger))

	// custom recovery middleware
	r.Use(gin.CustomRecovery(lib.Recover))

	// start server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			return
		}
	}()

	// graceful shutdown
	lib.GracefulShutdown(srv)
}
