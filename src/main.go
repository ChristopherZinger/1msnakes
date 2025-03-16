package main

import (
	"1msnakes/game"
	"1msnakes/vectors"
	"log"
	"net/http"
)

func main() {
	api()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func api() {
	_game := game.InitSnakeGame()
	playerMgr := game.NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		snake := game.CreateSnake([]*vectors.Vector{{X: 0, Y: 0}, {X: 0, Y: 10}})

		connection, err := playerMgr.CreateWebsocketConnection(w, r)
		if err != nil {
			log.Println("failed_to_create_connection")
			return
		}

		channel := make(chan game.GameEvent)
		player := game.CreatePlayer(
			snake,
			connection,
			channel,
			playerMgr,
		)

		_game.AddPlayer(player)
	})
}
