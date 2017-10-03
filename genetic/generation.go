package genetic

import "math/rand"

type generation struct {
	genes  [1000]*gene
	best   [4]*gene
	goal   [8]int
	Avgfit float64
}

func NewGeneration(goal [8]int) *generation {
	p := generation{}
	for i := 0; i < len(p.genes); i++ {
		p.genes[i] = newGene()
	}
	for i := 0; i < len(p.best); i++ {
		p.best[i] = newGene()
	}
	p.goal = goal
	p.Avgfit = 0
	return &p
}

func (p generation) GetBest() [4]*gene {
	return p.best
}

func (p *generation) Choice() {
	for i := 0; i < len(p.genes); i++ {
		for j := 0; j < len(p.best); j++ {
			if p.genes[i].getFitness(p.goal) > p.best[j].getFitness(p.goal) {
				p.best[j] = p.genes[i]
				break
			}
		}
	}
	p.Avgfit = 0
	for i := 0; i < len(p.genes); i++ {
		p.Avgfit += float64(p.genes[i].getFitness(p.goal))
	}
	p.Avgfit /= 1000
}

func (p generation) Crossover() *generation {
	new := NewGeneration(p.goal)
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

func (p generation) Mutent() {
	for i := 0; i < len(p.genes); i++ {
		p.genes[i].mutent()
	}
}
