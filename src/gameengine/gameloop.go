package gameengine

import (
	"1msnakes/player"
	"log"
	"time"
)

type GameState struct {
	players []*player.Player
}

func (g *GameState) AddPlayer(players *player.Player) {
	g.players = append(g.players, players)
}

func theLoop(g *GameState) {
	previousLoopTime := time.Now()
	i := 0
	for {
		if time.Since(previousLoopTime).Milliseconds() > 100 {
			log.Println("Game Loop: ", i)
			i++
			// update
			g.applySnakeMoves()

			// redraw
			for _, p := range g.players {
				ev := player.Event{
					Type: "hello",
				}
				p.Channel <- ev
			}
			previousLoopTime = time.Now()
		}
	}
}

func (g *GameState) applySnakeMoves() {
	for _, player := range g.players {
		player.Snake.ApplyNextMove()
	}
}

func InitSnakeGame() *GameState {
	log.Println("Init Snake Game!")

	var game = &GameState{}
	go theLoop(game)
	return game
}
