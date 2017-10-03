package main

import (
	"fmt"

	"github.com/khoneki/GeneticBitString/genetic"
)

func main() {
	var goal = [8]int{0, 0, 0, 0, 1, 1, 1, 1}
	parent := genetic.NewGeneration(goal)
	son := genetic.NewGeneration(goal)
	for parent.avgfit < 8 {
		parent.choice()
		fmt.Print(parent.avgfit)
		son = parent.crossover()
		son.mutent()
		parent = son
	}
}
