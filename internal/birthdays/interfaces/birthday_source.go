package interfaces

import (
	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
)

type BirthdaySource interface {
	GetBirthdays() ([]*entities.Birthday, error)
}
