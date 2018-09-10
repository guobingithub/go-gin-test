package main

import "go-gin-test/routers"

func main() {
	router := routers.InitRouter()

	// Start serving the application
	router.Run()
}
