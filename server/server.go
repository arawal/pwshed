package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/arawal/pwshed/hashlib"
	"github.com/arawal/pwshed/logger"
	"github.com/arawal/pwshed/stats"
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
			logger.Error("", fmt.Sprintf("listen: %s\n", err))
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
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/")
	{
		api.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
			logger.Info("/", fmt.Sprintf("%d", http.StatusOK))
		})

		api.POST("/hash", func(c *gin.Context) {
			startTime := time.Now()
			td := 5
			var err error
			td, err = strconv.Atoi(c.Query("timer"))
			if err != nil {
				td = 5
			}
			timer := time.NewTimer(time.Duration(td) * time.Second)

			password := c.PostForm("password")

			// optional: allow algorithm choice
			alg := c.Query("alg")

			result, err := hashlib.Hash(password, alg)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				logger.Error("/hash", err.Error())
				return
			}

			<-timer.C
			c.String(http.StatusOK, result)
			stats.UpdateStats(float64(time.Since(startTime)) / float64(time.Millisecond))
			logger.Info("/hash", fmt.Sprintf("passwords hashed this session: %d", stats.CurrentStats.Count))
		})

		api.GET("/shutdown", func(c *gin.Context) {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})

		api.GET("/stats", func(c *gin.Context) {
			data, err := stats.GetCurrentStats()
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				logger.Error("/stats", err.Error())
				return
			}
			c.JSON(http.StatusOK, data)
			logger.Info("/stats", "returned current stats")
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(404, "Page not found")
		logger.Error(c.Request.URL.Path, "path not found")
	})
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
	logger.Info("/shutdown", "received shutdown request")

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
	log.Println("Server gracefully exited")
	stats.UpdateStatsInStore()
}
