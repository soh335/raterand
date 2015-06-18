package raterand

import (
	"math/rand"
	"testing"
)

func TestRateRand(t *testing.T) {
	r1 := 0
	r2 := 0

	r := NewRateRand()
	r.Add(Choice{10, func() { r1++ }})
	r.Add(Choice{20, func() { r2++ }})

	g := r.Generate()

	loop := 100000
	for i := 0; i < loop; i++ {
		g().(func())()
	}

	if e, g := loop, r1+r2; e != g {
		t.Errorf("should %v got %v", e, g)
	}

	if !(2.0*0.95 <= float64(r2)/float64(r1) && float64(r2)/float64(r1) <= 2.0*1.05) {
		t.Errorf("r2/r1 about 20/10")
	}
}

func TestRateRandWithRand(t *testing.T) {
	r1 := 0
	r2 := 0

	r := NewRateRand()
	r.Rand = rand.New(rand.NewSource(1))
	r.Add(Choice{10, func() { r1++ }})
	r.Add(Choice{20, func() { r2++ }})

	g := r.Generate()

	loop := 100000
	for i := 0; i < loop; i++ {
		g().(func())()
	}

	if e, g := loop, r1+r2; e != g {
		t.Errorf("should %v got %v", e, g)
	}

	if !(2.0*0.95 <= float64(r2)/float64(r1) && float64(r2)/float64(r1) <= 2.0*1.05) {
		t.Errorf("r2/r1 about 20/10")
	}
}
