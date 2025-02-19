package main

import (
	"context"
	"flag"
	"log"

	"github.com/dxtym/webfetch/internal/server"
)

func main() {
	art := flag.String("art", "web/assets/art.txt", "art to use")
	port := flag.String("port", "6969", "port to use")

	flag.Parse()

	var ctx context.Context
	ctx = context.Background()
	ctx = context.WithValue(ctx, "art", *art)
	ctx = context.WithValue(ctx, "port", *port)

	if err := server.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
