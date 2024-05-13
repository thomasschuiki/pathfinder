//go:generate stringer -type=Race,Ability -output enum_string.go
package pathfinder

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root        = filepath.Join(filepath.Dir(b), "../..")
	ProjectRoot = filepath.Join(filepath.Dir(b), "..")
)

type Storage interface {
	Save()
	Load() interface{}
}

type Character struct {
	Name          string
	Ancestry      AncestryFormat
	Modifiers     map[Ability]int
	AbilityScores AbilityScores
	storage       Storage
}

var ancestryPath = filepath.Join(ProjectRoot, "refscraper/output/ancestry_clean.json")

func NewChar(name string, race Race, abilityScores AbilityScores, s Storage) *Character {
	char := &Character{
		Name:          name,
		Ancestry:      *NewAncestry(race),
		Modifiers:     make(map[Ability]int),
		AbilityScores: abilityScores,
		storage:       s,
	}

	char.CalcModifiers()

	return char
}

// func (c *Character) String() string {
// 	return fmt.Sprintf("%s the %s", c.Name, c.Race)
// }

func (c *Character) CalcModifiers() {
	for ab := Strength; ab <= Charisma; ab++ {
		c.AbilityScores[ab] = 10
		c.Modifiers[ab] = int(math.Floor(float64(c.AbilityScores[ab])/2.0) - 5.0)
		if slices.Contains(c.Ancestry.Source.Attribute, ab.String()) {
			c.Modifiers[ab] += 2
		}
		if slices.Contains(c.Ancestry.Source.AttributeFlaw, ab.String()) {
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

func NewAncestry(p Race) *AncestryFormat {

	ancestryFromFile := loadAncestry(p)
	return &ancestryFromFile
}

func loadAncestry(p Race) AncestryFormat {
	filter := func(data *AncestryFormat) bool {
		if data.Source.Name == p.String() {
			return true
		} else {
			return false
		}
	}

	ancestries := readJSON(ancestryPath, filter)

	return *ancestries[0]
}

func readJSON(fileName string, filter func(*AncestryFormat) bool) Ancestries {
	datas := Ancestries{}

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &datas)
	if err != nil {
		panic(err)
	}

	filteredData := Ancestries{}

	for _, data := range datas {
		// Do some filtering
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

type Ancestries []*AncestryFormat

type AncestryFormat struct {
	Index  string  `json:"_index"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source struct {
		SourceRaw []string `json:"source_raw"`
		Hp        int      `json:"hp"`
		Language  []string `json:"language"`
		Source    []string `json:"source"`
		Type      string   `json:"type"`
		Speed     struct {
			Max  int `json:"max"`
			Land int `json:"land"`
		} `json:"speed"`
		Weakness struct {
		} `json:"weakness"`
		SizeID            []int    `json:"size_id"`
		Trait             []string `json:"trait"`
		Attribute         []string `json:"attribute"`
		ID                string   `json:"id"`
		Text              string   `json:"text"`
		ExcludeFromSearch bool     `json:"exclude_from_search"`
		SkillMod          struct {
		} `json:"skill_mod"`
		Summary        string   `json:"summary"`
		Image          []string `json:"image"`
		TraitGroup     []string `json:"trait_group"`
		SourceCategory string   `json:"source_category"`
		RemasterID     string   `json:"remaster_id"`
		AttributeFlaw  []string `json:"attribute_flaw"`
		Resistance     struct {
		} `json:"resistance"`
		SpeedRaw    string   `json:"speed_raw"`
		URL         string   `json:"url"`
		Vision      string   `json:"vision"`
		Size        []string `json:"size"`
		ReleaseDate string   `json:"release_date"`
		RarityID    int      `json:"rarity_id"`
		TraitRaw    []string `json:"trait_raw"`
		Name        string   `json:"name"`
		Category    string   `json:"category"`
		HpRaw       string   `json:"hp_raw"`
		Rarity      string   `json:"rarity"`
	} `json:"_source"`
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
