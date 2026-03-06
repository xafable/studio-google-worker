package adapters

import (
	"time"

	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormBirthdayRepo struct {
	DB *gorm.DB
}

func NewGormBirthdayRepo(db *gorm.DB) *GormBirthdayRepo {
	return &GormBirthdayRepo{
		DB: db,
	}
}

func (g GormBirthdayRepo) SaveBirthdays(birthdays []*entities.Birthday) error {
	return g.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"date"}),
	}).Create(birthdays).Error
}

func (g GormBirthdayRepo) FindByBirthdayDate(date time.Time) ([]*entities.Birthday, error) {
	var b []*entities.Birthday

	g.DB.
		Where(
			"EXTRACT(MONTH FROM date) = ? AND EXTRACT(DAY FROM date) = ?",
			int(date.Month()),
			date.Day(),
		).
		Find(&b)

	return b, nil
}

func (g GormBirthdayRepo) SaveRecipient(recip entities.Recipient) error {
	return g.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "contact_type"}, {Name: "contact_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Create(&recip).Error
}

func (g GormBirthdayRepo) GetRecipientsByType(t entities.ContactType) ([]*entities.Recipient, error) {
	var r []*entities.Recipient
	err := g.DB.Where("contact_type = ?", t).Find(&r).Error

	return r, err
}
