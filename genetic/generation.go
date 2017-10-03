package genetic

import "math/rand"

type generation struct {
	genes  [1000]*gene
	best   [4]*gene
	goal   [8]int
	avgfit float64
}

func newGeneration(goal [8]int) *generation {
	p := generation{}
	var inv [8]int
	for i := 0; i < len(goal); i++ {
		if goal[i] == 0 {
			inv[i] = 1
		} else {
			inv[i] = 0
		}
	}
	for i := 0; i < len(p.genes); i++ {
		p.genes[i] = newGene()
	}
	for i := 0; i < len(p.best); i++ {
		p.best[i] = newGene()
		p.best[i].bit = inv
	}
	p.goal = goal
	return &p
}

func (p *generation) choice() {
	for i := 0; i < len(p.genes); i++ {
		for j := 0; j < len(p.best); j++ {
			if p.genes[i].getFitness(p.goal) == p.best[j].getFitness(p.goal) {
				p.best[j] = p.genes[i]
				break
			}
		}
	}
	for i := 0; i < len(p.best); i++ {
		p.avgfit += float64(p.best[i].getFitness(p.goal))
	}
	p.avgfit /= 4
}

func (p generation) crossover() *generation {
	new := newGeneration(p.goal)
	for i := 0; i < len(p.genes); i++ {
		p1 := rand.Intn(4)
		p2 := rand.Intn(4)
		a := rand.Intn(8)
		for j := 0; j < a; j++ {
			new.genes[i].bit[j] = p.best[p1].bit[j]
		}
		for j := a; j < 8; j++ {
			new.genes[i].bit[j] = p.best[p2].bit[j]
		}
	}
	return new
}

func (p generation) mutent() {
	for i := 0; i < len(p.genes); i++ {
		p.genes[i].mutent()
	}
}
