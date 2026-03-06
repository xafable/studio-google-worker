package interfaces

import (
	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
)

type Sender interface {
	Send(message entities.SenderMessage) error
	GetType() entities.ContactType
}
