package vectors

import (
	"fmt"
	"math"
)

type Directions int

const (
	VecW Directions = iota
	VecE
	VecN
	VecS
)

var DirToVec = map[Directions]Vector{
	VecW: {X: -1, Y: 0},
	VecE: {X: 1, Y: 0},
	VecS: {X: 0, Y: -1},
	VecN: {X: 0, Y: 1},
}

type Vector struct{ X, Y float64 }

func (v *Vector) Len() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v *Vector) Unit() *Vector {
	vResult := &Vector{
		X: float64(v.X) / v.Len(),
		Y: float64(v.Y) / v.Len(),
	}
	return vResult
}

func (v *Vector) print() {
	fmt.Printf("Vector x: %f, y: %f. \n", v.X, v.Y)
}

func (self *Vector) IsEqual(vec *Vector) bool {
	return self.X == vec.X && self.Y == vec.Y
}

func VectorSum(vectors []*Vector) *Vector {
	vector := Vector{}
	for _, v := range vectors {
		vector.X += v.X
		vector.Y += v.Y
	}
	return &vector
}

func VectorSubstract(v1, v2 *Vector) *Vector {
	v := Vector{X: v1.X - v2.X, Y: v1.Y - v2.Y}
	return &v
}

func AreVectorsParallel(v1, v2 *Vector) bool {
	return v1.X*v2.Y == v1.Y*v2.X
}

func DoVectorsShareDirection(v1, v2 *Vector) bool {
	dotProd := v1.X*v2.X + v1.Y*v2.Y
	cos := float64(dotProd) / (math.Sqrt(float64(v1.X*v1.X+v1.Y*v1.Y) * float64(v2.X*v2.X+v2.Y*v2.Y)))
	result := cos == float64(1)

	return result
}
