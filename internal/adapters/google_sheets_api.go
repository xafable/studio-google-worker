package adapters

import (
	"fmt"
	"time"

	"github.com/xafable/studio-google-worker/internal/entities"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"context"
)

type GoogleSheetsApi struct {
	SheetID string `json:"sheet_id"`
}

func NewSheetsApi(sheetID string) *GoogleSheetsApi {
	return &GoogleSheetsApi{
		SheetID: sheetID,
	}
}

func (g GoogleSheetsApi) GetBirthdays() ([]*entities.Birthday, error) {

	ctx := context.Background()

	creds, err := google.FindDefaultCredentials(ctx, sheets.SpreadsheetsScope)
	if err != nil {
		fmt.Printf("error when get greds: %s", err)
		return nil, err
	}

	srv, err := sheets.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		fmt.Printf("error when create sheets service: %s", err)
		return nil, err
	}

	range_ := "page1!B:C"
	resp, err := srv.Spreadsheets.Values.Get(g.SheetID, range_).Do()
	if err != nil {
		fmt.Printf("error when get sheets values: %s", err)
		return nil, err
	}

	fmt.Println("resp sheets")
	var result []*entities.Birthday

	for _, row := range resp.Values {
		name := row[0]
		date := row[1]
		t, err := time.Parse("02.01.2006", date.(string))
		if err != nil {
			continue
		}

		result = append(result, &entities.Birthday{
			Name: name.(string),
			Date: t,
		})
		fmt.Println(name, t.Format("2 January 2006"))

	}

	return result, nil
}
