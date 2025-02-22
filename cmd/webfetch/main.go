package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/dxtym/webfetch/internal/server"
	"github.com/dxtym/webfetch/internal/utils"
)

func main() {
	var (
		art  = flag.String("art", "web/assets/art.txt", "ascii art")
		port = flag.String("port", "6969", "port number")
		help = flag.Bool("help", false, "help message")
	)

	flag.Parse()

	if *help {
		utils.ShowHelp()
		os.Exit(0)
	}

	var ctx context.Context
	ctx = context.Background()
	ctx = context.WithValue(ctx, "art", *art)
	ctx = context.WithValue(ctx, "port", *port)

	if err := server.Run(ctx); err != nil {
		log.Fatalf("failed to run app: %s", err)
	}
}
