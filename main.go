package main

import (
	"MyGram/database"
	"MyGram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	err := r.Run()
	if err != nil {
		return
	}
}
