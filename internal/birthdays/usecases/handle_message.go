package usecases

import (
	"strconv"

	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
	"github.com/xafable/studio-google-worker/internal/birthdays/interfaces"
)

func HandleMessage(pu entities.PollerMessage, rep interfaces.BirthdayRepository) {
	chatID, _ := strconv.ParseInt(pu.From, 10, 64)

	recip := entities.Recipient{
		Name:        pu.Name,
		ContactID:   strconv.FormatInt(chatID, 10),
		ContactType: pu.Type,
	}

	rep.SaveRecipient(recip)
}
