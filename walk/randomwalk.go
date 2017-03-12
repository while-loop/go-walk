package walk

import (
	"math/rand"
	"time"
	"errors"
)

// Weight of a direction
type Weight uint32

type weights *[4]Weight

// Functions to execute when a direction is chosen
type Walker interface {
	Left()
	Right()
	Up()
	Down()
}

const (
	// No need to make these public
	lEFT  = 0
	rIGHT = 1
	uP    = 2
	dOWN  = 3
)

var r *rand.Rand // your private random

func init() {
	// set up the random seed for our code
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Random Walk configuration
type RandomWalk struct {
	Left, Right, Up, Down Weight
	Walker                Walker
}

// Create new Random Walk with custom probabilities for each direction
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
	if w.Walker == nil {
		return errors.New("Walker implemenation not set")
	}

	ws, sum := w.preprocess()
	if sum <= 0 {
		return errors.New("Sum of weights is <= 0")
	}

	for i := uint32(0); i < iterations; i++ {
		switch getRandy(ws, sum) {
		case lEFT:
			w.Walker.Left()
			break
		case rIGHT:
			w.Walker.Right()
			break
		case uP:
			w.Walker.Up()
			break
		case dOWN:
			w.Walker.Down()
			break
		}
	}
	return nil
}

// Get a random direction given a set of normalized weights
// Using the Sum of Weights method
func getRandy(ws weights, sum uint64) int {
	if sum <= 0 {
		return -1
	}

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
