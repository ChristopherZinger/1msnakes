package player

import (
	"1msnakes/snake"
	"1msnakes/vectors"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type GameEvent struct {
	Type string
}

type WebsocketMsg struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type Player struct {
	Snake   *snake.Snake
	Channel chan GameEvent

	manager    *Manager
	connection *websocket.Conn
}

type PlayerList map[*Player]bool

func CreatePlayer(snake *snake.Snake, connection *websocket.Conn, channel chan GameEvent, mgr *Manager) *Player {
	player := Player{
		Snake:      snake,
		connection: connection,
		Channel:    channel,
		manager:    mgr,
	}

	go player.writeMessage()
	go player.receiveMesssages()

	return &player
}

func (p *Player) writeMessage() {
	for {
		select {
		case _, ok := <-p.Channel:
			if !ok {
				err := p.connection.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			snakePixels, err := json.Marshal(p.Snake.GetPixels())
			if err != nil {
				log.Println("could not convert snake pixels to json")
			}

			data, err := json.Marshal(WebsocketMsg{
				Type:    "snake_position",
				Payload: snakePixels,
			})
			if err != nil {
				log.Println("failed to marshal", err)
				return
			}

			if err := p.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("failed to send message, %v", err)
			}
		}
	}
}

type SnakeMoveMessage struct {
	Direction int
}

func (p *Player) receiveMesssages() {
	defer func() {
		p.manager.RemovePlayer(p)
	}()

	for {
		_, payload, err := p.connection.ReadMessage()
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

		var request WebsocketMsg
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println("failed to unmarshall the event")
			break
		}

		var move SnakeMoveMessage
		if err := json.Unmarshal(request.Payload, &move); err != nil {
			log.Println("failed to unmarshal move event:", err)
			break
		}

		switch dir := move.Direction; dir {
		case 0:
			p.Snake.SetNextMv(vectors.VecN)
		case 1:
			p.Snake.SetNextMv(vectors.VecN)
		case 2:
			p.Snake.SetNextMv(vectors.VecN)
		case 3:
			p.Snake.SetNextMv(vectors.VecN)
		default:
			p.Snake.SetNextMv(vectors.VecN)
		}
	}
}