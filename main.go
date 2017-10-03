package main

import (
	"fmt"

	"github.com/khoneki/GeneticBitString/genetic"
)

func main() {
	var goal = [8]int{0, 0, 0, 0, 1, 1, 1, 1}
	parent := genetic.NewGeneration(goal)
	son := genetic.NewGeneration(goal)
	parent.Choice()
	for i := 0; parent.Avgfit < 8; i++ {
		parent.Choice()
		fmt.Print(i, " ", parent.Avgfit, "\n")
		son = parent.Crossover()
		son.Mutent()
		parent = son
	}
}
