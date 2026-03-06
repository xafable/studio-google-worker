package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
	"gorm.io/gorm"
)

func CreateBirthdaysTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20260113_create_birthdays",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&entities.Birthday{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("birthdays")
		},
	}
}
