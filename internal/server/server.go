package server

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dxtym/webfetch/internal/socket"
)

type Param struct {
	Art string
}

func Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	ws := socket.NewWebSocket()
	http.HandleFunc("/update", ws.Update)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("web/views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		file, err := os.ReadFile("web/assets/art.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		t.Execute(w, Param{Art: string(file)})
	})

	go func() {
		port := fmt.Sprintf(":%v", ctx.Value("port"))

		log.Printf("server listening on: %s\n", port)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatalf("failed to listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}
