package interfaces

import (
	"github.com/xafable/studio-google-worker/internal/entities"
)

type BirthdaySource interface {
	GetBirthdays() ([]*entities.Birthday, error)
}
