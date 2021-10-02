package handlers

import (
	"net/http"

	"blipblop/src/storage"

	"github.com/gin-gonic/gin"
)

func GetJobs(c *gin.Context) {
	jobs, err := storage.DB().GetJobs()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, jobs)
}
