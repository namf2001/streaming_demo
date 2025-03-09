package models

import (
	"time"

	"gorm.io/gorm"
)

// Record is a struct that represents a record in the database
type Record struct {
	gorm.Model

	ID       int `gorm:"primary_key"`
	Keyword  string
	Node     string
	CreateAd time.Time
	UpdateAp time.Time
}
