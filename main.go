package main

import (
	"MyGram/database"
	"MyGram/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := os.Getenv("PORT")
	port = "3000"
	r.Run(":" + port)
}
