package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dxtym/webfetch/internal/socket"
)

func Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	ws := socket.NewWebSocket()

	http.Handle("/", http.FileServer(http.Dir("web/views")))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.HandleFunc("/update", ws.Update)

	go func() {
		log.Println("server listening on: :6969")
		err := http.ListenAndServe(":6969", nil)
		if err != nil {
			log.Fatalf("failed to listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}
