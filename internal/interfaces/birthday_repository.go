package interfaces

import (
	"time"

	"github.com/xafable/studio-google-worker/internal/entities"
)

type BirthdayRepository interface {
	SaveBirthdays(birthdays []*entities.Birthday)
	FindByBirthdayDate(date time.Time) ([]*entities.Birthday, error)
}
