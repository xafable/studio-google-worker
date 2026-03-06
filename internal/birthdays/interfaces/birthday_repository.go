package interfaces

import (
	"time"

	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
)

type BirthdayRepository interface {
	SaveBirthdays(birthdays []*entities.Birthday) error
	FindByBirthdayDate(date time.Time) ([]*entities.Birthday, error)
	SaveRecipient(recip entities.Recipient) error
	GetRecipientsByType(t entities.ContactType) ([]*entities.Recipient, error)
}
