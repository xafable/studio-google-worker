package usecases

import (
	"fmt"

	"github.com/xafable/studio-google-worker/internal/birthdays/adapters"
	"github.com/xafable/studio-google-worker/internal/birthdays/interfaces"
)

type BirthdaysJob struct {
	Name               string
	SheetID            string
	BirthdayRepository interfaces.BirthdayRepository
}

func NewBirthdaysJob(name string, sheetID string, birthdayRepository interfaces.BirthdayRepository) *BirthdaysJob {
	return &BirthdaysJob{
		Name:               name,
		SheetID:            sheetID,
		BirthdayRepository: birthdayRepository,
	}
}

func (b BirthdaysJob) Do() error {
	sapi := adapters.NewSheetsApi(b.SheetID)
	gb := NewBirthdaysServ(b.BirthdayRepository, sapi)

	bd, err := gb.Get()
	if err != nil {
		fmt.Println("error when Get in bjob")
		return err
	}

	gb.Save(bd)

	return nil
}
