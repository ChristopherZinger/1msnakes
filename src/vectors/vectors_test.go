package vectors

import (
	"fmt"
	"testing"
)

func TestDoVectorsShareDirection(t *testing.T) {
	type Case struct {
		v1, v2 Vector
		result bool
	}

	cases := [...]Case{
		{
			Vector{X: 20, Y: 0},
			Vector{X: 10, Y: 0},
			true,
		},
		{
			Vector{X: 5, Y: 5},
			Vector{X: 2, Y: 2},
			true,
		},
		{
			Vector{X: 0, Y: 5},
			Vector{X: 0, Y: -5},
			false,
		},
		{
			Vector{X: 5, Y: 5},
			Vector{X: -5, Y: -5},
			false,
		},
		{
			Vector{X: 5, Y: 5},
			Vector{X: 5, Y: -5},
			false,
		},
		{
			Vector{X: 11, Y: 7},
			Vector{X: 33, Y: 21},
			true,
		},
	}

	for _, c := range cases[:] {
		if DoVectorsShareDirection(&c.v1, &c.v2) != c.result {
			fmt.Printf("expected %t for x1: %f, y1: %f AND x2: %f, y2: %f.\n", c.result, c.v1.X, c.v1.Y, c.v2.X, c.v2.Y)

			t.Error("failed same direction test")
		}
	}

}

func TestAreVectorsParallel(t *testing.T) {
	type Case struct {
		v1, v2 Vector
		result bool
	}

	cases := [...]Case{
		{
			Vector{0, 5},
			Vector{0, -5},
			true,
		},
		{
			Vector{5, 0},
			Vector{-5, 0},
			true,
		},
		{
			Vector{0, 5},
			Vector{0, -5},
			true,
		},
		{
			Vector{1, 5},
			Vector{-1, -5},
			true,
		},
		{
			Vector{1, 5},
			Vector{2, 10},
			true,
		},
		{
			Vector{1, 0},
			Vector{0, 1},
			false,
		},
	}

	for _, c := range cases[:] {
		if AreVectorsParallel(&c.v1, &c.v2) != c.result {
			fmt.Printf(
				"Expected AreVectorsParallel v1x:%f v1y:%f AND v2x:%f v2y:%f to be %t",
				c.v1.X, c.v1.Y, c.v2.X, c.v2.Y, c.result,
			)
			t.Error(`failed AreVectorsParallel`)
		}
	}
}

func TestVectorSum(t *testing.T) {
	type Case struct {
		vectors []*Vector
		result  Vector
	}

	cases := [...]Case{
		{
			[]*Vector{{1, 1}, {1, 1}},
			Vector{2, 2},
		},
		{
			[]*Vector{{2, 2}, {-1, -1}},
			Vector{1, 1},
		},
		{
			[]*Vector{{2, 2}, {-1, -1}, {10, 5}},
			Vector{11, 6},
		},
	}

	for _, c := range cases {
		r := VectorSum(c.vectors)
		if c.result.X != r.X || c.result.Y != r.Y {
			fmt.Printf("Expected Result: x: %f, y: %f.\n", c.result.X, c.result.Y)
			fmt.Printf("Received Result: x: %f, y: %f.\n", r.X, r.Y)
			t.Error("Failed Summing Vectors")
		}

	}
}
