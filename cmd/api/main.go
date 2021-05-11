package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/stnnnghm/go-api-design/pkg/application"
	"github.com/stnnnghm/go-api-design/pkg/exithandler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
