package entities

import (
	"time"
)

type Birthday struct {
	ID   uint      `gorm:"primaryKey"`
	Name string    `gorm:"uniqueIndex"`
	Date time.Time `gorm:"type:date"`
}
