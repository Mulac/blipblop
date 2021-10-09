package main

import (
	"blipblop/src/core/config"
	"fmt"
)

func main() {
	fmt.Printf("running core with config %+v\n", config.Config())
}
