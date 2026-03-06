package usecases

import (
	"fmt"
	"time"

	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
	"github.com/xafable/studio-google-worker/internal/birthdays/interfaces"
)

type SendBirthdaysJob struct {
	Name               string
	Sender             interfaces.Sender
	BirthdayRepository interfaces.BirthdayRepository
}

func NewSendBirthdaysJob(name string, sender interfaces.Sender, birthdayRepository interfaces.BirthdayRepository) *SendBirthdaysJob {
	return &SendBirthdaysJob{
		Name:               name,
		Sender:             sender,
		BirthdayRepository: birthdayRepository,
	}
}

func (sb SendBirthdaysJob) Do() error {
	recipients, _ := sb.BirthdayRepository.GetRecipientsByType(sb.Sender.GetType())
	text := "Сьогодні в команді Stefhani Studio – важлива подія!💫 "

	now := time.Now()
	d := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	birthdays, err := sb.BirthdayRepository.FindByBirthdayDate(d)
	if err != nil {
		fmt.Println("error when get FindByBirthdayDate", err)
		return err
	}

	if len(birthdays) == 0 {
		fmt.Println("no birthdays")
		return nil
	}

	for _, bi := range birthdays {
		text += fmt.Sprintf("\n\nДень народження у %s ", bi.Name)
	}

	for _, r := range recipients {
		sb.Sender.Send(entities.SenderMessage{
			To:   r.ContactID,
			Text: text,
		})
	}
	return nil
}
