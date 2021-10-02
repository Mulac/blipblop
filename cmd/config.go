package main

import (
	. "blipblop/core/config"
	"fmt"
)

func main() {
	fmt.Printf("running core with config %+v\n", Config())
}
