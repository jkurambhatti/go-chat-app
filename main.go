package main

import (
	"log"

	"github.com/auth0/auth0-golang/examples/regular-web-app/app"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Init()
	go globalRoom.run()
	StartServer()

}
