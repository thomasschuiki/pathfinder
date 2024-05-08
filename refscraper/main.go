package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"gopkg.in/yaml.v2"
    "pathfinder-tui/refscraper/googleapis"
)

var version = "0.0.0"

// Config ...
type Config struct {
	Refsheets struct {
		SpellSheetID      string `yaml:"spellSheetID"`
		FeatsSheetID      string `yaml:"featsSheetID"`
		MagicItemsSheetID string `yaml:"magicItemsSheetID"`
		BestiarySheetID   string `yaml:"bestiarySheetID"`
	} `yaml:"Refsheets"`
}

func main() {
	cfgPath := flag.String("c", "./config.yml", "path to the config.yml file")
	fVersion := flag.Bool("v", false, "output version")
	flag.Parse()

	if *fVersion {
		fmt.Printf("Version: %v\n", version)
	}

	var cfg Config
	getConfigFile(*cfgPath, &cfg)

	outputPath := "./output"

	// gSheetsService := googleapis.NewGoogleClient("./credentials.json")
	// downloadAsCSV(cfg.Refsheets.SpellSheetID,"spells.csv", *gSheetsService)
	// downloadAsCSV(cfg.Refsheets.BestiarySheetID,"bestiary.csv", *gSheetsService)
	// downloadAsCSV(cfg.Refsheets.FeatsSheetID,"feats.csv", *gSheetsService)
	// downloadAsCSV(cfg.Refsheets.MagicItemsSheetID,"magicitems.csv", *gSheetsService)
	categories := []string{
		"action",
		"ancestry",
		"archetype",
		"armor",
		"article",
		"background",
		"class",
		"creature",
		"creature-family",
		"deity",
		"equipment",
		"feat",
		"hazard",
		"rules",
		"skill",
		"shield",
		"spell",
		"source",
		"trait",
		"weapon",
		"weapon-group",
	}
	escfg := elasticsearch.Config{
		Addresses: []string{
			"https://elasticsearch.aonprd.com/",
		},
	}
	es, err := elasticsearch.NewClient(escfg)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range categories {
		q := getQuery(v)
		fmt.Println(q)
		res, err := es.Search(
			es.Search.WithIndex("aon"),
			es.Search.WithBody(strings.NewReader(q)),
			es.Search.WithFrom(0),
			es.Search.WithSize(10000))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		b, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(filepath.Join(outputPath, fmt.Sprintf("%s.json", v)), b, 0o664)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getQuery(cat string) string {
	return fmt.Sprintf(`{
    "query": {
      "match": {
        "category": "%s"
      }
    },
    "_source": {
      "exclude": [
        "*markdown*",
        "navigation"
      ]
    }
  }`, cat)
}

func getConfigFile(path string, cfg *Config) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		log.Fatal(err)
	}
}

// GetRangeFormatted gets formatted values in 'a1Range' from the spreadsheet
// doc identified by 'id'.
// All values are returned as strings, formatted as they display in the spreadsheet document
func convertToSliceOfStrings(v [][]interface{}) ([][]string, error) {
	// make a [][]string to hold typecast values
	stringVals := make([][]string, len(v))
	for i := range stringVals {
		stringVals[i] = make([]string, len(v[i]))
	}
	// cast each interface{} to a string
	for r, row := range v {
		for c, v := range row {
			stringVals[r][c] = v.(string)
		}
	}
	return stringVals, nil
}

func downloadAsCSV(sheetID, filename string, gSheetsService googleapis.GoogleClient) {
	sheetInfo, err := gSheetsService.GetSheetInfo(sheetID)
	if err != nil {
		log.Fatalln(err)
	}
	spells, err := gSheetsService.GetSheetData(sheetID, sheetInfo.Sheets[0].Properties.Title, "ROWS")
	if err != nil {
		log.Fatalln(err)
	}

	outfile, err := os.Create(filename)
	defer outfile.Close()

	if err != nil {
		log.Fatalln("failed to create csv file", err)
	}

	stringValues, err := convertToSliceOfStrings(spells.Values)
	if err != nil {
		log.Fatalln("could not convert to slice of strings", err)
	}

	w := csv.NewWriter(outfile)
	err = w.WriteAll(stringValues)
	if err != nil {
		log.Fatalln("failed to write to csv file", err)
	}
}
