package game

import (
	"1msnakes/vectors"
	"fmt"
	"testing"
)

func TestSnakeMove(t *testing.T) {
	type Case struct {
		mvDir     vectors.Directions
		newHead   vectors.Vector
		newTail   vectors.Vector
		snake     []*vectors.Vector
		numPoints int
	}

	cases := [...]Case{
		{
			mvDir:     vectors.VecE,
			newHead:   vectors.Vector{X: 11, Y: 0},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 2,
		},
		{
			mvDir:     vectors.VecN,
			newHead:   vectors.Vector{X: 10, Y: 1},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 3,
		},
		{
			mvDir:     vectors.VecS,
			newHead:   vectors.Vector{X: 10, Y: -1},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 3,
		},
		{
			mvDir:     vectors.VecW,
			newHead:   vectors.Vector{X: 9, Y: 0},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 0, Y: 0}},
			newTail:   vectors.Vector{X: 1, Y: 0},
			numPoints: 3,
		},
		{
			mvDir:     vectors.VecE,
			newHead:   vectors.Vector{X: 11, Y: 0},
			snake:     []*vectors.Vector{{X: 10, Y: 0}, {X: 5, Y: 0}, {X: 5, Y: 1}},
			newTail:   vectors.Vector{X: 5, Y: 0},
			numPoints: 2,
		},
	}

	for _, c := range cases {
		s := CreateSnake(c.snake)

		s.Move(c.mvDir)

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

func TestGetSnakePixels(t *testing.T) {

	type Case struct {
		snake  []*vectors.Vector
		pixels []*vectors.Vector
	}

	cases := [...]Case{
		{
			snake: []*vectors.Vector{{X: 4, Y: 0}, {X: 0, Y: 0}},
			pixels: []*vectors.Vector{
				{X: 4, Y: 0},
				{X: 3, Y: 0},
				{X: 2, Y: 0},
				{X: 1, Y: 0},
				{X: 0, Y: 0},
			},
		},
		{
			snake: []*vectors.Vector{{X: 4, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 2}},
			pixels: []*vectors.Vector{
				{X: 4, Y: 0},
				{X: 3, Y: 0},
				{X: 2, Y: 0},
				{X: 2, Y: 1},
				{X: 2, Y: 2},
			},
		},
	}

	for _, c := range cases {
		s := CreateSnake(c.snake)
		pixels := s.GetPixels()

		if len(pixels) != len(c.pixels) {
			t.Error("Incorrect number of pixels")
		}

		for i, px := range pixels {
			if px.X != c.pixels[i].X || px.Y != c.pixels[i].Y {
				t.Error("Incorrect snake pixel positions")
			}

		}
	}

}
