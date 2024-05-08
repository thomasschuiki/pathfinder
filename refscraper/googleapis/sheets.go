package googleapis

import (
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

// GetSheetData gets data from a google sheet
func (g *GoogleClient) GetSheetData(spreadsheetId, readRange, majorDimension string) (*sheets.ValueRange, error) {
	resp, err := g.service.Spreadsheets.Values.Get(spreadsheetId, readRange).MajorDimension(majorDimension).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	return resp, nil
}

// GetSheetInfo gets info from a google sheet
func (g *GoogleClient) GetSheetInfo(spreadsheetId string) (*sheets.Spreadsheet, error) {
	resp, err := g.service.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve info from sheet: %v", err)
	}

	if len(resp.Sheets) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	return resp, nil
}
