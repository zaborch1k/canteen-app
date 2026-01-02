package main

import (
	"log"

	"canteen-app/internal/app"
)

//	@title			CanteenApp API
//	@version		1.0
//	@description	Internal API for canteen web app
//	@license.name	Apache 2.0
//	@host			localhost:8080

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("starting server on :8080")
	if err := a.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
