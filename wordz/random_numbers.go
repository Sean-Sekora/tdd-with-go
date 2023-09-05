package wordz

import (
	"math/rand"
	"time"
)

type RandomNumbers interface {
	NextInt(int) int
}

type RandomNumbersImpl struct {
}

func (RandomNumbersImpl) NextInt(upperBound int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(upperBound)
}
