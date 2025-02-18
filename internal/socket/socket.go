package socket

import (
	"log"
	"net/http"
	"time"

	"github.com/dxtym/webfetch/internal/specs"
	"github.com/gorilla/websocket"
)

type WebSocket struct {
	websocket.Upgrader
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (ws *WebSocket) Update(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("failed to upgrade: %s\n", err)
		return
	}
	defer conn.Close()
	
	for {
		host, err := specs.GetHostInfo()
		if err != nil {
			log.Printf("cannot obtain host: %s\n", err)
			break
		}

		mem, err := specs.GetMemInfo()
		if err != nil {
			log.Printf("cannot obtain mem: %s\n", err)
			break
		}
		
		cpu, err := specs.GetCpuInfo()
		if err != nil {
			log.Printf("cannot obtain cpu: %s\n", err)
			break
		}

		msg := []byte(host + cpu + mem)
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("failed to write: %s\n", err)
			break
		}

		time.Sleep(time.Second * 3)
	}
}
