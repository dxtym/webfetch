package server

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dxtym/webfetch/internal/socket"
)

type Param struct {
	Art string
}

//go:embed web/css
var css embed.FS

//go:embed web/views
var html embed.FS

//go:embed web/assets/art.txt
var txt embed.FS

func Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	ws := socket.NewWebSocket()
	http.HandleFunc("/update", ws.Update)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.FS(css))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFS(html, "web/views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var file []byte
		art := ctx.Value("art").(string)
		file, err = txt.ReadFile(art)
		if err != nil {
			if _, ok := err.(*fs.PathError); !ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			file, err = os.ReadFile(art)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
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
