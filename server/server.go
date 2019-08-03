package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arawal/pwshed/hashlib"
	"github.com/gin-gonic/gin"
)

// LaunchServer configures and launches a http server
func LaunchServer() {
	// Start and run the server
	router := initRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	handleGracefulShutdown(srv)
}

// initRouter initializes a gin router with the preconfigured routes
/*
	Input:
		- none
	Output:
		- router - *gin.Eingine - gin router
*/
func initRouter() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/")
	{
		api.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		api.POST("/hash", func(c *gin.Context) {
			timer := time.NewTimer(5 * time.Second)
			password := c.PostForm("password")

			// optional: allow algorithm choice
			alg := c.Query("alg")

			result, err := hashlib.Hash(password, alg)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			<-timer.C
			c.String(http.StatusOK, result)
		})

		api.GET("/shutdown", func(c *gin.Context) {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})
	}
	return router
}

func handleGracefulShutdown(srv *http.Server) {

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down server:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server gracefully exiting")
}
