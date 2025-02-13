package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dxtym/xifetch/internal/specs"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	http.Handle("/", http.FileServer(http.Dir("web/views")))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.HandleFunc("/ws", handleUpdate)

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

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade: %s\n", err)
		return
	}
	defer conn.Close()

	for {
		host, err := specs.GetHostInfo()
		if err != nil {
			log.Printf("cannot obtain host: %s\n", err)
			return
		}

		cpu, err := specs.GetCpuInfo()
		if err != nil {
			log.Printf("cannot obtain cpu: %s\n", err)
			return
		}

		mem, err := specs.GetMemInfo()
		if err != nil {
			log.Printf("cannot obtain mem: %s\n", err)
			return
		}

		msg := []byte(host + cpu + mem)
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("failed to write: %s\n", err)
			return
		}

		time.Sleep(time.Second * 3)
	}
}
