package main

import (
	"first-api/Config"
	"first-api/Routes"
	"log"
)

func main() {
	Config.InitDB() // Initialize the database connection

	// Your code here

	router := Routes.SetupRouter()

	// Running
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
