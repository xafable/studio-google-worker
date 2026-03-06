package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
	"gorm.io/gorm"
)

func CreateRecipientsTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20260114_create_recipients",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&entities.Recipient{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("recipients")
		},
	}
}
