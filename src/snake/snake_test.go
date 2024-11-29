package snake

import (
	"1msnakes/vectors"
	"fmt"
	"testing"
)

func TestSnakeMove(t *testing.T) {
	type Case struct {
		mv        vectors.Vector
		newHead   vectors.Vector
		numPoints int
	}

	cases := [...]Case{
		{
			mv:        vectors.Vector{X: 1, Y: 0},
			newHead:   vectors.Vector{X: 11, Y: 0},
			numPoints: 2,
		},
		{
			mv:        vectors.Vector{X: 0, Y: 1},
			newHead:   vectors.Vector{X: 10, Y: 1},
			numPoints: 3,
		},
		{
			mv:        vectors.Vector{X: 0, Y: -1},
			newHead:   vectors.Vector{X: 10, Y: -1},
			numPoints: 3,
		},
		{
			mv:        vectors.Vector{X: -1, Y: 0},
			newHead:   vectors.Vector{X: 9, Y: 0},
			numPoints: 3,
		},
	}

	for _, c := range cases {
		s := CreateSnake()

		s.Move(&c.mv)

		head := s.getHead()

		if head.X != c.newHead.X || head.Y != c.newHead.Y || len(s.Body) != c.numPoints {
			fmt.Printf("expected:\t x: %d, y: %d, num: %d\n", c.newHead.X, c.newHead.Y, c.numPoints)
			fmt.Printf("headX:\t\t x: %d, y: %d, num: %d\n", head.X, head.Y, len(s.Body))
			fmt.Println("----")
			t.Error("failed to move snake head\n")
		}
	}

}
