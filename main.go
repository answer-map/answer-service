package main

import (
	"log"

	"github.com/answer-map/answer-service/app"
)

func main() {
	config, err := app.ConfigFromEnvironmentVariables()
	if err != nil {
		log.Fatalf("failed to create config, error = %v", err)
	}

	app, err := app.NewApp(config)
	if err != nil {
		log.Fatalf("failed to create app, error = %v", err)
	}

	app.Run()
}
