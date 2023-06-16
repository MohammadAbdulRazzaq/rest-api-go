package main

import (
	"first-api/Routes"
	"first-api/docker/db"
	"log"
)

func init() {

	db.ReturnDB()
}
func main() {
	//Config.InitDB() // Initialize the database connection

	router := Routes.SetupRouter()

	// Running
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
