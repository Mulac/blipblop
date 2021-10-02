package backend

import (
	"blipblop/src/backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// We can use groups for routes with common paths and middleware
	app := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "blipblop", // user: password
	}))

	app.GET("jobs", handlers.GetJobs)

	return r
}
