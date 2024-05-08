//go:generate stringer -type=Race,Ability -output enum_string.go
package pathfinder

import (
	"math"
	"slices"
)

type Storage interface {
	Save(interface{})
	Load(interface{})
}

type Character struct {
	Name          string
	Race          Race
	Ancestry      Ancestry
	Modifiers     map[Ability]int
	AbilityScores AbilityScores
	storage       Storage
}

func NewChar(name string, race Race, abilityScores AbilityScores, storage Storage) *Character {
	char := &Character{
		Name:          name,
		Race:          race,
		Modifiers:     make(map[Ability]int),
		AbilityScores: abilityScores,
		storage:       storage,
	}

	char.CalcModifiers()

	return char
}

// func (c *Character) String() string {
// 	return fmt.Sprintf("%s the %s", c.Name, c.Race)
// }

func (c *Character) CalcModifiers() {
	a := NewAncestry(c.Race)
	for ab := Strength; ab <= Charisma; ab++ {
		c.AbilityScores[ab] = 10
		c.Modifiers[ab] = int(math.Floor(float64(c.AbilityScores[ab])/2.0) - 5.0)
		if slices.Contains(a.Bonuses, ab) {
			c.Modifiers[ab] += 2
		}
		if slices.Contains(a.Flaws, ab) {
			c.Modifiers[ab] -= 2
		}
	}
}

type AbilityScores map[Ability]int

func NewAbilityScores(str, dex, con, int, wis, cha int) AbilityScores {
	abs := make(AbilityScores)
	abs[Strength] = str
	abs[Dexterity] = dex
	abs[Constitution] = con
	abs[Intelligence] = int
	abs[Wisdom] = wis
	abs[Charisma] = cha
	return abs
}

type Race int

const (
	Dwarf Race = iota
	Elf
	Gnome
	HalfElf
	HalfOrc
	Halfling
	Human
)

type Ancestry struct {
	Race    Race
	Bonuses []Ability
	Flaws   []Ability
}

func NewAncestry(p Race) *Ancestry {
	// load ancestry from storage
	// use dwarf for now
	// bonus: str, wis
	// flaw: cha
	return &Ancestry{Race: p, Bonuses: []Ability{Strength, Wisdom}, Flaws: []Ability{Charisma}}
}

type Ability int

const (
	Strength Ability = iota
	Dexterity
	Constitution
	Intelligence
	Wisdom
	Charisma
)
