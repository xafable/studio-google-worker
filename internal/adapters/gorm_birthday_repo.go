package adapters

import (
	"time"

	"github.com/xafable/studio-google-worker/internal/entities"
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

func (g GormBirthdayRepo) SaveBirthdays(birthdays []*entities.Birthday) {
	g.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"date"}),
	}).Create(birthdays)
}

func (g GormBirthdayRepo) FindByBirthdayDate(date time.Time) ([]*entities.Birthday, error) {
	var b []*entities.Birthday
	g.DB.Where("date = ?", date).Find(&b)

	return b, nil
}
