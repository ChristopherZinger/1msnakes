package snake

import (
	"1msnakes/arrays"
	"1msnakes/vectors"
)

type Snake struct{ Body []*vectors.Vector }

type Directions int8

const (
	W Directions = iota
	E
	N
	S
)

// it only work for mv vectors of length = 1
// becase tail always move by 1 unit
func (s *Snake) Move(mv *vectors.Vector) {
	currentHead := s.getHead()
	// currentTail := s.getTail()

	newHead := vectors.VectorSum((&[...]*vectors.Vector{currentHead, mv})[:])

	if vectors.DoVectorsShareDirection(s.getHeadVec(), mv) {
		currentHead.X = newHead.X
		currentHead.Y = newHead.Y
	} else {
		//vLen := currentTail.Len()
		// uv := vectors.Vector{
		//	X: float64(currentTail.X) / vLen,
		//	Y: float64(currentTail.Y) / vLen,
		//}

		s.Body = arrays.Prepend(s.Body, newHead)
	}

	// TODO: move the tail

	tailVec := s.getTailVec()
	if tailVec.Len() > 1 {
		// is tail vector > 1 ? move tail point
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

func (s *Snake) getHead() *vectors.Vector {
	return (s.Body)[0]
}

func (s *Snake) getTail() *vectors.Vector {
	return (s.Body)[len(s.Body)-1]
}

func (s *Snake) getHeadVec() *vectors.Vector {
	v1, v2 := (s.Body)[0], (s.Body)[1]
	return vectors.VectorSubstract(v1, v2)
}

// TODO: create test
func (s *Snake) getTailVec() *vectors.Vector {
	bodyLen := len(s.Body)
	v1, v2 := (s.Body)[bodyLen-1], (s.Body)[bodyLen-2]
	return vectors.VectorSubstract(v2, v1)
}

// TODO: pass initial shape or points
func CreateSnake() *Snake {
	arr := [10]*vectors.Vector{}
	body := arr[:2]
	body[0] = &vectors.Vector{X: 10, Y: 0}
	body[1] = &vectors.Vector{X: 0, Y: 0}

	s := Snake{Body: body}
	return &s
}
