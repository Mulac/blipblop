package main

import (
	. "core/config"
	"fmt"
)

func main() {
	fmt.Printf("running core with config %+v\n", Config())
}
