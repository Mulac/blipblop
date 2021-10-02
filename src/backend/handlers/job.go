package handlers

import (
	"net/http"
	"sync"

	"blipblop/src/storage"

	"github.com/gin-gonic/gin"
)

var db storage.JobDatabase
var once sync.Once

func GetJobs(c *gin.Context) {
	once.Do(func() {
		db = storage.NewJobDatabase()
	})

	jobs, err := db.GetJobs()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, jobs)
}
