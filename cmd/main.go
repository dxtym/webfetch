package main

import (
	"context"
	"log"

	"github.com/dxtym/minefetch/internal/server"
)

func main() {
	ctx := context.Background()

	if err := server.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
