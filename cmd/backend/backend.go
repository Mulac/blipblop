package main

import "blipblop/src/backend"

func main() {
	r := backend.SetupRouter()
	r.Run("")
}
