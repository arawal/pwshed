package server

import (
	"net/http"
	"time"

	"github.com/arawal/pwshed/hashlib"
	"github.com/gin-gonic/gin"
)

// LaunchServer configures and launches a http server
func LaunchServer() {
	// Start and run the server
	initRouter().Run(":8080")
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
	}
	return router
}
