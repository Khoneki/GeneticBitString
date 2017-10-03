package genetic

import "fmt"

func main() {
	var goal = [8]int{0, 0, 0, 0, 1, 1, 1, 1}
	parent := newGeneration(goal)
	son := newGeneration(goal)
	for parent.avgfit < 8 {
		parent.choice()
		fmt.Print(parent.avgfit)
		son = parent.crossover()
		son.mutent()
		parent = son
	}
}
