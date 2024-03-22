package main

import (
	"MyGram/database"
	"MyGram/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env is not loaded properly")
	}

	database.StartDB()
	r := router.StartApp()
	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		return
	}
}
