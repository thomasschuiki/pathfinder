package main

import (
	"fmt"

	pf "pathfinder-tui/pathfinder"
)

func main() {
    abilityScores := pf.NewAbilityScores(10, 10, 10, 10, 10, 10)
	char := pf.NewChar("Ragnulf", pf.Dwarf, abilityScores, nil)
	fmt.Printf("%#v", char)
}
