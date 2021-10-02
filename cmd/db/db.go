package main

import (
	"blipblop/src/storage"
	"fmt"
)

func main() {
	job, err := storage.DB().GetJobs()
	if err != nil {
		panic(err)
	}

	fmt.Println(job)
}
