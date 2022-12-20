package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type HealthCheck struct {
	Component string `json:"component" binding:"required"`
	Status    string `json:"status"`
}

/**
* Main Function
 */
func main() {
	// Gin default Router
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	healthCheck := router.Group("/healthCheck")
	{
		healthCheck.GET("/middleware", func(c *gin.Context) {
			result := callMiddleware("middleware")
			c.JSON(200, result)
		})
	}

	// Start & Run server
	router.Run(":3000")
}

func callMiddleware(component string) string {
	client := resty.New()
	var h HealthCheck
	hostName := os.Getenv("MW_HOST_NAME")
	hostPort := os.Getenv("MW_HOST_PORT")

	var middlewareUrl string = "http://" + hostName + ":" + hostPort + "/healthCheck"
	log.Println(middlewareUrl)
	_, err := client.R().
		SetBody(HealthCheck{Component: component}).
		SetResult(&h).
		Post(middlewareUrl)
	if err != nil {
		log.Printf("Front End: unable to connect Middleware (%v)", middlewareUrl)
		return "error"
	}

	log.Printf("Front End: Middleware status: %v", h.Status)
	return h.Status
}
