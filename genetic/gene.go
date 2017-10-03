package genetic

import (
	"math/rand"
)

type gene struct {
	bit [8]int
}

func newGene() *gene {
	g := gene{}
	for i := 0; i < 8; i++ {
		g.bit[i] = rand.Intn(2)
	}
	return &g
}

func (g gene) getFitness(goal [8]int) (fit int) {
	fit = 0
	for i := 0; i < 8; i++ {
		if g.bit[i] == goal[i] {
			fit++
		}
	}
	return
}

func (g *gene) mutent() {
	a := rand.Intn(8)
	b := rand.Float64()
	if b < 0.01 {
		if g.bit[a] == 0 {
			g.bit[a] = 1
		} else {
			g.bit[a] = 0
		}
	}
}
