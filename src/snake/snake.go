package snake

import (
	"1msnakes/arrays"
	"1msnakes/vectors"
)

type Snake struct{ Body []*vectors.Vector }

func (s *Snake) Move(d vectors.Directions) {
	mv := vectors.DirToVec[d]
	s.moveHead(&mv)
	s.moveTail()
}

func (s *Snake) GetPixels() []vectors.Vector {
	var pixels []vectors.Vector

	prevPixel := s.getHead()
	pixels = append(pixels, *prevPixel) // this should copy head ??

	for i := 0; i < len(s.Body)-1; i++ {
		mv := vectors.VectorSubstract(s.Body[i+1], s.Body[i]).Unit()
		for s.Body[i+1].IsEqual(prevPixel) == false {
			pixel := vectors.VectorSum([]*vectors.Vector{mv, prevPixel})
			pixels = append(pixels, *pixel)
			prevPixel = pixel
		}
	}
	return pixels
}

func (s *Snake) moveHead(mv *vectors.Vector) {
	currentHead := s.getHead()
	newHead := vectors.VectorSum((&[...]*vectors.Vector{currentHead, mv})[:])

	if vectors.DoVectorsShareDirection(s.getHeadVec(), mv) {
		currentHead.X = newHead.X
		currentHead.Y = newHead.Y
	} else {
		s.Body = arrays.Prepend(s.Body, newHead)
	}
}

func (s *Snake) moveTail() {
	tailVec := s.getTailVec()
	if tailVec.Len() > 1 {
		moveUnitVec := tailVec.Unit()
		newTail := vectors.VectorSum([]*vectors.Vector{moveUnitVec, s.getTail()})
		tailIndex := len(s.Body) - 1
		s.Body[tailIndex] = newTail
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

func CreateSnake(body []*vectors.Vector) *Snake {
	s := Snake{Body: body}
	return &s
}
