package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJobs(c *gin.Context) {
	fmt.Println("hello")
	c.Status(http.StatusOK)
}
