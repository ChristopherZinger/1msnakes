package game

import (
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
	Players PlayerList
	sync.RWMutex
}

func NewManager() *Manager {
	mgr := &Manager{
		Players: make(PlayerList),
	}

	return mgr
}

func (m *Manager) CreateWebsocketConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	log.Println("new connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func (m *Manager) addPlayer(player *Player) {
	m.Lock()
	defer m.Unlock()
	m.Players[player] = true
}

func (m *Manager) removePlayer(player *Player) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Players[player]; ok {
		player.connection.Close()
		delete(m.Players, player)
	}
}
