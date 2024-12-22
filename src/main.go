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
	manager := player.NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// create snake
		snake := snake.CreateSnake([]*vectors.Vector{{X: 0, Y: 0}, {X: 0, Y: 10}})

		// create connection
		connection, err := manager.CreateWebsocketConnection(w, r)
		if err != nil {
			log.Println("failed_to_create_connection")
			return
		}
		// create channel
		channel := make(chan player.Event)

		// create player (snake, connection, channel)
		player := player.CreatePlayer(
			snake,
			connection,
			channel,
			manager,
		)

		// addPlayerToGame(player)
		game.AddPlayer(player)

		go player.WriteMessage() // sendDataToUser
		// go player.receiveMessage(chanel) // setNextMove

	})
}
