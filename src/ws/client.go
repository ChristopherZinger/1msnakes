package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type WsClientList map[*WsClient]bool

type WsClient struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan Event
}

func NewClient(c *websocket.Conn, m *Manager) *WsClient {
	return &WsClient{
		connection: c,
		manager:    m,
		egress:     make(chan Event),
	}
}

func (c *WsClient) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		select {
		case messages, ok := <-c.egress:
			if !ok {
				err := c.connection.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			data, err := json.Marshal(messages)
			if err != nil {
				log.Println()
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("failed to send message, %v", err)
			}
		}
	}
}

func (c *WsClient) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println("failed to unmarshall the event")
			break
		}

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println(err)
		}
	}

}
