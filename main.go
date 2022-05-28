package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/answer-map/answer-service/app"
)

func main() {
	configAsBytes, err := os.ReadFile(`config/local.json`)
	if err != nil {
		log.Fatalf("failed to load config, error = %v", err)
	}

	var config app.Config
	if err := json.Unmarshal(configAsBytes, &config); err != nil {
		log.Fatalf("failed to unmarshal config, error = %v", err)
	}

	app, err := app.NewApp(&config)
	if err != nil {
		log.Fatalf("failed to create app, error = %v", err)
	}

	app.Run()
}
