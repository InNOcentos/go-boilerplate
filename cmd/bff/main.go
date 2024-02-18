package main

import (
	"context"
	"log"

	"github.com/InNOcentos/Doggo-track/backend/internal/bff/app"
)

// @title                                 Doggo-track API
// @version                               1.0
// @description                           Doggo-track API.

// @securityDefinitions.apiKey            ApiKeyAuth
// @in                                    header
// @name                                  Authorization

func main() {
	ctx := context.Background()
	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
