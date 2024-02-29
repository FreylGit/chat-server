package main

import (
	"context"
	app "github.com/FreylGit/chat-server/internal/app"
	"log"
)

func main() {
	ctx := context.Background()
	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed setup settigs server")
	}
	err = app.Run()
	if err != nil {
		log.Fatalf("Failed run server")
	}

}
