package main

import (
	"chat-server/internal/app"
	"chat-server/internal/config"
	"context"
	"log"
)

func main() {

	ctx := context.Background()
	err := config.Load("local.env")
	a, err := app.NewApp(ctx)

	if err != nil {
		log.Fatal(err)
	}
	a.Run(ctx)
}
