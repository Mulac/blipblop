package backend

import (
	"blipblop/src/backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	// TODO - make more secure
	corsConfig.AllowOrigins = []string{"*"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))

	// We can use groups for routes with common paths and middleware
	app := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "blipblop", // user: password
	}))

	app.GET("jobs", handlers.GetJobs)

	return r
}
