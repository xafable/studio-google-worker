package entities

type ContactType string

const (
	ContactTypeTelegram ContactType = "telegram"
	ContactTypeViber    ContactType = "viber"
)

type Recipient struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	ContactType ContactType `gorm:"uniqueIndex:uidx_contact"`
	ContactID   string      `gorm:"uniqueIndex:uidx_contact"`
}
