package main

import (
	"context"
	"log"

	"github.com/dxtym/xifetch/internal/server"
)

func main() {
	ctx := context.Background()

	if err := server.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
