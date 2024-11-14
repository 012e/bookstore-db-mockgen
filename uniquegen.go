package main

import "github.com/brianvoe/gofakeit/v7"

type UniqueNumGen struct {
	Min       int
	Max       int
	generated map[int]bool
}

// NewUniqueNumGen creates a new UniqueNumGen
func NewUniqueNumGen(min int, max int) *UniqueNumGen {
	return &UniqueNumGen{
		Min:       min,
		Max:       max,
		generated: make(map[int]bool),
	}
}

// Get generates a unique random number between min and max
func (g *UniqueNumGen) Get() int {
	if len(g.generated) == g.Max-g.Min+1 {
		panic("No more unique numbers to generate")
	}

	for {
		num := gofakeit.Number(g.Min, g.Max)
		if !g.generated[num] {
			g.generated[num] = true
			return num
		}
	}
}
