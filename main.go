package main

import (
	"MyGram/database"
	"MyGram/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env is not loaded properly")
	}

	database.StartDB()
	r := router.StartApp()
	err = r.Run()
	if err != nil {
		return
	}
}
