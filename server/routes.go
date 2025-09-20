package server

import (
	"encoding/json"
	"github.com/cobbles/ant-simulator/app"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}

const tickMs = 100

func registerRoutes(world *app.WorldStruct, mux *http.ServeMux) {
	mux.HandleFunc("/ws", wsHandler(world))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
}

func wsHandler(world *app.WorldStruct) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade error", err)
			return
		}
		defer conn.Close()

		for {
			time.Sleep(tickMs * time.Millisecond)

			data, _ := json.Marshal(world)

			conn.WriteMessage(websocket.TextMessage, data)
		}
	}
}
