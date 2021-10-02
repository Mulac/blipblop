package main

import (
	"db/db"
	"fmt"
)

func main() {
	db := db.NewJobDatabase()
	job, err := db.GetJob(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(job)
}
