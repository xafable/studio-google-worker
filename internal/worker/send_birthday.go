package worker

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xafable/studio-google-worker/internal/entities"
	"github.com/xafable/studio-google-worker/internal/interfaces"
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
	recipients := [1]int64{433380489}
	text := "Сьогодні в команді Stefhani Studio – важлива подія!💫 "

	d := time.Date(2002, 12, 7, 0, 0, 0, 0, time.UTC)
	birthdays, err := sb.BirthdayRepository.FindByBirthdayDate(d)
	if err != nil {
		fmt.Println("error when get FindByBirthdayDate", err)
		return err
	}

	for _, bi := range birthdays {
		text += fmt.Sprintf("\n\nДень народження у %s ", bi.Name)
	}

	for _, r := range recipients {
		sb.Sender.Send(entities.SenderMessage{
			To:   strconv.FormatInt(r, 10),
			Text: text,
		})
	}
	return nil
}
