package main

import (
	"api/routes"
)

func main() {
	// Get Routes Index
	e:= routes.Index()

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}