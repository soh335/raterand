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

func NewRateRand() *RateRand {
	return &RateRand{}
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

	randomFloat := rand.Float64
	if r.Rand != nil {
		randomFloat = r.Rand.Float64
	}

	return func() Item {
		index := randomFloat() * max
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
