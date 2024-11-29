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
		newTail   vectors.Vector
		snake     []*vectors.Vector
		numPoints int
	}

	cases := [...]Case{
		{
			mv:        vectors.Vector{X: 1, Y: 0},
			newHead:   vectors.Vector{X: 11, Y: 0},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 2,
		},
		{
			mv:        vectors.Vector{X: 0, Y: 1},
			newHead:   vectors.Vector{X: 10, Y: 1},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 3,
		},
		{
			mv:        vectors.Vector{X: 0, Y: -1},
			newHead:   vectors.Vector{X: 10, Y: -1},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 3,
		},
		{
			mv:        vectors.Vector{X: -1, Y: 0},
			newHead:   vectors.Vector{X: 9, Y: 0},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 3,
		},
		{
			mv:        vectors.Vector{X: 1, Y: 0},
			newHead:   vectors.Vector{X: 11, Y: 0},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 5, Y: 0}, {X: 5, Y: 1}},
			newTail:   vectors.Vector{X: 5, Y: 0},
			numPoints: 2,
		},
	}

	for _, c := range cases {
		s := CreateSnake(c.snake)

		s.Move(&c.mv)

		if len(s.Body) != c.numPoints {
			fmt.Printf("expected num of body points:\t : %d, but got: %d\n", c.numPoints, len(s.Body))
			t.Error("Wrong number of snake segments after move")
		}

		head := s.getHead()
		if head.X != c.newHead.X || head.Y != c.newHead.Y {
			fmt.Printf("expected head at:\t x: %f, y: %f\n", c.newHead.X, c.newHead.Y)
			fmt.Printf("got head at:\t x: %f, y: %f\n", head.X, head.Y)
			t.Error("Wrong snake head position after move")
		}

		tail := s.getTail()
		if tail.X != c.newTail.X || tail.Y != c.newTail.Y {
			fmt.Printf("expected tail at:\t x: %f, y: %f\n", c.newTail.X, c.newTail.Y)
			fmt.Printf("got tail at:\t x: %f, y: %f\n", tail.X, tail.Y)
			t.Error("Wrong snake tail position after move")
		}
	}
}
