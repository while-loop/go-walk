package walk

import (
	"math/rand"
	"time"
	"errors"
)

type Weight uint32

type weights *[4]Weight

type Walker interface {
	Left()
	Right()
	Up()
	Down()
}

const (
	lEFT  = 0
	rIGHT = 1
	uP    = 2
	dOWN  = 3
)

var r *rand.Rand // your private random

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type RandomWalk struct {
	Left, Right, Up, Down Weight
	walker                Walker
}

func NewRandomWalk(left, right, up, down Weight, walker Walker) *RandomWalk {
	return &RandomWalk{left, right, up, down, walker}
}

// Create new Random Walk with equal probability for all directions
func New(walker Walker) *RandomWalk {
	return &RandomWalk{1, 1, 1, 1, walker}
}

// [0] = Left
// [1] = Right
// [2] = Up
// [3] = Down
func (w RandomWalk) preprocess() (weights, uint64) {
	return &[4]Weight{w.Left, w.Right, w.Up, w.Down}, uint64(w.Left) + uint64(w.Right) + uint64(w.Up) + uint64(w.Down)
}

// Perform a random walk with the amount of iterations given
// Calls directions from the Walker interface given
//
// Error is return if walker implementation was not set
func (w *RandomWalk) Walk(iterations uint32) error {
	if w.walker == nil {
		return errors.New("Walker implemenation not set")
	}

	ws, sum := w.preprocess()
	for i := uint32(0); i < iterations; i++ {

		switch getRandy(ws, sum) {
		case lEFT:
			w.walker.Left()
			break
		case rIGHT:
			w.walker.Right()
			break
		case uP:
			w.walker.Up()
			break
		case dOWN:
			w.walker.Down()
			break
		}
	}
	return nil
}

// Get a random direction given a set of normalized weights
// Using the Sum of Weights method
func getRandy(ws weights, sum uint64) int {
	randy := r.Int63n(int64(sum))
	ttl := int64(0)

	for j := 0; j < len(ws); j++ {
		ttl += int64(ws[j])
		if randy < ttl {
			return j
		}
	}
	return -1
}