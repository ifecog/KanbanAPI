package models

import (
	"time"

	"github.com/google/uuid"
)

type Courier struct {
	ID			uint		`gorm:"primaryKey"`
	UUID		uuid.UUID		`gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	FirstName	string		`gorm:"size:100;not null"`
	LastName	string		`gorm:"size:100;not null"`
	Email		string		`gorm:"size:100;not null;unique"`
	Phone		string		`gorm:"size:15;not null"`
	Password	string		`gorm:"not null"`
	IsActive	bool		`gorm:"default:true"`
	CreatedAt	time.Time	`gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime"`
}


