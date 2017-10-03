package main

import (
	"fmt"

	"github.com/khoneki/GeneticBitString/genetic"
)

func main() {
	var goal = [8]int{0, 1, 0, 0, 1, 0, 1, 1}
	parent := genetic.NewGeneration(goal)
	son := genetic.NewGeneration(goal)
	for i := 0; parent.Avgfit < 8; i++ {
		parent.Choice()
		fmt.Print(i, " ", parent.GetBest()[0], " ", parent.Avgfit, "\n")
		if parent.Avgfit >= 7.9 {
			break
		}
		son = parent.Crossover()
		son.Mutent()
		parent = son
	}
}
