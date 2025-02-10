package server

import (
	"log"
	"net/http"
	"time"

	"github.com/dxtym/zfetch/internal/specs"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Run() {
	http.Handle("/", http.FileServer(http.Dir("web/views/")))
	http.HandleFunc("/ws", handleUpdate)

	log.Fatal(http.ListenAndServe(":6969", nil))
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}
	defer conn.Close()

	for {
		host, err := specs.GetHostInfo()
		if err != nil {
			log.Println("host: ", err)
			return
		}

		cpu, err := specs.GetCpuInfo()
		if err != nil {
			log.Println("cpu: ", err)
			return
		}

		mem, err := specs.GetMemInfo()
		if err != nil {
			log.Println("mem: ", err)
			return
		}

		msg := []byte(host + cpu + mem)
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("write: ", err)
			return
		}

		time.Sleep(time.Second * 5)
	}
}
