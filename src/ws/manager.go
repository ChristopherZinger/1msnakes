package ws

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients  WsClientList
	handlers map[string]EventHandler
	sync.RWMutex
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(WsClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessage
}

func (m *Manager) routeEvent(ev Event, c *WsClient) error {
	if handler, ok := m.handlers[ev.Type]; ok {
		if err := handler(ev, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("Unknown event type")
	}
}

func SendMessage(ev Event, c *WsClient) error {
	fmt.Println(ev)
	c.egress <- ev

	return nil

}

func (m *Manager) ServerWs(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)
	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addClient(c *WsClient) {
	m.Lock()
	defer m.Unlock()
	m.clients[c] = true
}

func (m *Manager) removeClient(c *WsClient) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[c]; ok {
		c.connection.Close()
		delete(m.clients, c)
	}
}
