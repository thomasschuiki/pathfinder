package googleapis

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleClient struct {
	service *sheets.Service
}

func NewGoogleClient(jsonPath string) *GoogleClient {
	ctx := context.Background()
	gSheetSvc, err := sheets.NewService(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return &GoogleClient{service: gSheetSvc}
}
