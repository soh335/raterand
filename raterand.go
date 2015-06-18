package raterand

// porting from Sub::Rate::NoMaxRate
// https://github.com/typester/Sub-Rate/blob/master/lib/Sub/Rate/NoMaxRate.pm

import (
	"math/rand"
)

type Item interface{}

type Choice struct {
	Weight float64
	Item   Item
}

func NewRateRand(r *rand.Rand) *RateRand {
	return &RateRand{
		Rand: r,
	}
}

type RateRand struct {
	choices []Choice
	Rand    *rand.Rand
}

func (r *RateRand) Add(c Choice) {
	r.choices = append(r.choices, c)
}

func (r *RateRand) Generate() Generator {
	var max float64
	for _, choice := range r.choices {
		max += choice.Weight
	}

	choices := r.choices
	_r := r.Rand

	return func() Item {
		index := _r.Float64() * max
		var cursor float64

		for _, choice := range choices {
			cursor += choice.Weight

			if index <= cursor {
				return choice.Item
			}
		}

		return nil
	}
}

type Generator func() Item
