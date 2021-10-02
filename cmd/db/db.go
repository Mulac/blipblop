package main

import (
	"blipblop/src/storage"
	"fmt"
)

func main() {
	db := storage.NewJobDatabase()
	job, err := db.GetJobs()
	if err != nil {
		panic(err)
	}

	fmt.Println(job)
}
