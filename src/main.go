package main

import (
	"1msnakes/ws"
	"log"
	"net/http"
)

func main() {
	api()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func api() {
	manager := ws.NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))
	http.HandleFunc("/ws", manager.ServerWs)
}
