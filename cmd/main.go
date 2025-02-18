package main

import (
	"context"
	"flag"
	"log"

	"github.com/dxtym/webfetch/internal/server"
)

func main() {
	port := flag.String("port", "6969", "port to use")
	flag.Parse()
	
	c := context.Background()
	ctx := context.WithValue(c, "port", *port)

	if err := server.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
