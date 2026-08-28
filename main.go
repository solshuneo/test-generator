package main

import "test-generator/router"

func main() {
	router := router.SetupRouter()

	router.Run(":3001")
}
