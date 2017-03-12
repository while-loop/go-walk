package walk

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"
	"github.com/google/gofuzz"
	"log"
)

func TestConstructorSetsVals(t *testing.T) {
	l, r, u, d := genWeights()
	rw := NewRandomWalk(l, r, u, d, nil)
	assert.Equal(t, l, rw.Left)
	assert.Equal(t, r, rw.Right)
	assert.Equal(t, u, rw.Up)
	assert.Equal(t, d, rw.Down)
	assert.Nil(t, rw.Walker)
}

func TestInterfaceCallsMethods(t *testing.T) {

	maxIts := 100000
	l := Weight(0)
	r := Weight(0)
	u := Weight(100000)
	d := Weight(0)
	mw := &TestWalker{}
	rw := NewRandomWalk(l, r, u, d, mw)
	rw.Walk(uint32(maxIts))

	assert.Equal(t, 0, mw.l, "left sums not 0")
	assert.Equal(t, 0, mw.r, "right sums not 0")
	assert.Equal(t, maxIts, mw.u, "up sums not maxIts")
	assert.Equal(t, 0, mw.d, "down sums not 0")
}

func TestInterfaceCallsMethodsEqually(t *testing.T) {
	maxIts := 10000000
	thrshld := maxIts / 5.0
	l := Weight(1000)
	r := Weight(1000)
	u := Weight(1000)
	d := Weight(1000)
	mw := &TestWalker{}
	rw := NewRandomWalk(l, r, u, d, mw)
	rw.Walk(uint32(maxIts))

	assert.True(t, mw.l >= thrshld, "left sums not thrshld", mw.l)
	assert.True(t, mw.r >= thrshld, "right sums not thrshld", mw.r)
	assert.True(t, mw.u >= thrshld, "up sums not thrshld", mw.u)
	assert.True(t, mw.d >= thrshld, "down sums not thrshld", mw.d)

	sum := mw.l + mw.d + mw.r + mw.u
	assert.Equal(t, maxIts, sum, "not all iterations met", sum)
}

func TestNilWalkerReturnsError(t *testing.T) {
	assert.EqualError(t, New(nil).Walk(32), "Walker implemenation not set")
}

func TestMaxWeights(t *testing.T) {
	maxIts := 1000000
	thrshld := maxIts / 5.0
	w := Weight(4294967295)
	mw := &TestWalker{}
	rw := NewRandomWalk(w, w, w, w, mw)
	rw.Walk(uint32(maxIts))

	assert.True(t, mw.l >= thrshld, "left sums not thrshld", mw.l)
	assert.True(t, mw.r >= thrshld, "right sums not thrshld", mw.r)
	assert.True(t, mw.u >= thrshld, "up sums not thrshld", mw.u)
	assert.True(t, mw.d >= thrshld, "down sums not thrshld", mw.d)

	sum := mw.l + mw.d + mw.r + mw.u
	assert.Equal(t, maxIts, sum, "not all iterations met", sum)
}

func TestPreprocessSumsCorrectly(t *testing.T) {
	l, r, u, d := genWeights()
	_, sum := NewRandomWalk(l, r, u, d, nil).preprocess()
	assert.Equal(t, uint64(l)+uint64(r)+uint64(u)+uint64(d), sum, "", l, r, u, d)
}

func TestPreprocessAssignsCorrectDirEnum(t *testing.T) {
	l, r, u, d := genWeights()
	weights, sum := NewRandomWalk(l, r, u, d, nil).preprocess()

	assert.Equal(t, 4, len(weights))
	assert.Equal(t, l, weights[lEFT])
	assert.Equal(t, r, weights[rIGHT])
	assert.Equal(t, u, weights[uP])
	assert.Equal(t, d, weights[dOWN])
	assert.Equal(t, uint64(l)+uint64(r)+uint64(u)+uint64(d), sum, "", l, r, u, d)
}

func TestPrependedZeroWeightsDontExecute(t *testing.T) {
	maxIts := 100000
	sums := map[int]int{
		lEFT:  0,
		rIGHT: 0,
		uP:    0,
		dOWN:  0,
	}

	weights, sum := NewRandomWalk(0, 0, 0, 1, nil).preprocess()
	for i := 0; i < maxIts; i++ {
		sums[getRandy(weights, sum)]++
	}

	assert.Equal(t, 4, len(sums), "Unknown direction given", sums)

	assert.Equal(t, 0, sums[lEFT], "left sums not 0", sums[lEFT])
	assert.Equal(t, 0, sums[rIGHT], "right sums not 0", sums[rIGHT])
	assert.Equal(t, 0, sums[uP], "up sums not 0", sums[uP])
	assert.Equal(t, maxIts, sums[dOWN], "not all dirs were down", sums[dOWN])
}

func TestHalfWeightsDontExecute(t *testing.T) {
	maxIts := 1000000
	sums := map[int]int{
		lEFT:  0,
		rIGHT: 0,
		uP:    0,
		dOWN:  0,
	}

	weights, sum := NewRandomWalk(1, 0, 0, 1, nil).preprocess()
	for i := 0; i < maxIts; i++ {
		sums[getRandy(weights, sum)]++
	}

	assert.Equal(t, 4, len(sums), "Unknown direction given")

	assert.True(t, sums[lEFT] >= maxIts/3.0, "left sums not distributed", sums[lEFT], maxIts/3.0)
	assert.Equal(t, 0, sums[rIGHT], "right sums not 0", sums[rIGHT])
	assert.Equal(t, 0, sums[uP], "up sums not 0", sums[uP])
	assert.True(t, sums[dOWN] >= maxIts/3.0, "down sums not distributed", sums[dOWN], maxIts/3.0)
}

func TestZeroWeightsReturnsNegative1(t *testing.T) {
	maxIts := 100000
	sums := map[int]int{
		lEFT:  0,
		rIGHT: 0,
		uP:    0,
		dOWN:  0,
	}

	weights, sum := NewRandomWalk(0, 0, 0, 0, nil).preprocess()
	for i := 0; i < maxIts; i++ {
		sums[getRandy(weights, sum)]++
	}

	assert.Equal(t, 5, len(sums), "Unknown direction given", sums)

	assert.Equal(t, maxIts, sums[-1], "-1 sums not max Its", sums[-1])
	assert.Equal(t, 0, sums[lEFT], "left sums not 0", sums[lEFT])
	assert.Equal(t, 0, sums[rIGHT], "right sums not 0", sums[rIGHT])
	assert.Equal(t, 0, sums[uP], "up sums not 0", sums[uP])
	assert.Equal(t, 0, sums[dOWN], "down sums not 0", sums[dOWN])
}

func TestZeroWeightsReturnErrorWalking(t *testing.T) {
	maxIts := uint32(100000)

	mw := &TestWalker{}
	rw := NewRandomWalk(0, 0, 0, 0, mw)
	err := rw.Walk(maxIts)

	assert.EqualError(t, err, "Sum of weights is <= 0", err.Error())

	assert.Equal(t, 0, mw.l, "left sums not 0")
	assert.Equal(t, 0, mw.r, "right sums not 0")
	assert.Equal(t, 0, mw.u, "up sums not 0")
	assert.Equal(t, 0, mw.d, "down sums not 0")
}

func TestWalkWithFuzz(t *testing.T) {
	var maxIts uint32
	fuzz.New().NilChance(0).Funcs(func(e *uint32, c fuzz.Continue){
		*e = uint32(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(5000000) + 10000)
	}).Fuzz(&maxIts)

	mw := &TestWalker{}
	var rw RandomWalk
	fuzz.New().NilChance(0).Funcs(func(e *RandomWalk, c fuzz.Continue) {
		c.Fuzz(&e.Left)
		c.Fuzz(&e.Right)
		c.Fuzz(&e.Up)
		c.Fuzz(&e.Down)
	}).Fuzz(&rw)

	rw.Walker = mw

	err := rw.Walk(maxIts)
	assert.NoError(t, err, "", mw, rw)
	log.Println(mw, rw, maxIts)
}

type TestWalker struct {
	l, r, u, d int
}

func (w *TestWalker) Left() {
	w.l++
}

func (w *TestWalker) Right() {
	w.r++
}

func (w *TestWalker) Up() {
	w.u++
}

func (w *TestWalker) Down() {
	w.d++
}

func
genWeights() (Weight, Weight, Weight, Weight) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Weight(rng.Uint32()), Weight(rng.Uint32()), Weight(rng.Uint32()), Weight(rng.Uint32())
}
