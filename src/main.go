package main

import (
	"1msnakes/gameengine"
	"1msnakes/player"
	"1msnakes/snake"
	"1msnakes/vectors"
	"log"
	"net/http"
)

func main() {
	api()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func api() {
	game := gameengine.InitSnakeGame()
	playerMgr := player.NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		snake := snake.CreateSnake([]*vectors.Vector{{X: 0, Y: 0}, {X: 0, Y: 10}})

		connection, err := playerMgr.CreateWebsocketConnection(w, r)
		if err != nil {
			log.Println("failed_to_create_connection")
			return
		}

		channel := make(chan player.GameEvent)
		player := player.CreatePlayer(
			snake,
			connection,
			channel,
			playerMgr,
		)

		game.AddPlayer(player)
	})
}
