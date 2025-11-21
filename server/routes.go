package server

import (
	"encoding/json"
	"github.com/cobbles/ant-simulator/app"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	tickMs       = 100
	pingInterval = 30 * time.Second
	pongWait     = 35 * time.Second
	writeWait    = 10 * time.Second
)

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type hub struct {
	connections map[*connection]bool
	register    chan *connection
	unregister  chan *connection
	broadcast   chan []byte
	mutex       sync.RWMutex
}

var wsHub = &hub{
	connections: make(map[*connection]bool),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	broadcast:   make(chan []byte),
}

func registerRoutes(world *app.TWorld, mux *http.ServeMux) {
	go wsHub.run()
	go worldBroadcaster(world)

	mux.HandleFunc("/ws", wsHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.mutex.Lock()
			h.connections[c] = true
			h.mutex.Unlock()
			log.Println("WebSocket client connected")

		case c := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
				log.Println("WebSocket client disconnected")
			}
			h.mutex.Unlock()

		case message := <-h.broadcast:
			h.mutex.RLock()
			for c := range h.connections {
				select {
				case c.send <- message:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

func worldBroadcaster(world *app.TWorld) {
	ticker := time.NewTicker(tickMs * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := json.Marshal(world)
			if err != nil {
				log.Printf("Error marshaling world: %v", err)
				continue
			}
			wsHub.broadcast <- data
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error", err)
		return
	}

	c := &connection{
		ws:   conn,
		send: make(chan []byte, 256),
	}

	wsHub.register <- c

	go c.writePump()
	go c.readPump()
}

func (c *connection) readPump() {
	defer func() {
		wsHub.unregister <- c
		c.ws.Close()
	}()

	c.ws.SetReadLimit(512)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (c *connection) writePump() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.ws.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
