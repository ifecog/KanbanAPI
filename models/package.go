package models

import (
	"time"

	"github.com/google/uuid"
)

type PackageSize string


type Package struct {
	ID          		uint        `gorm:"primaryKey"`
	UUID        		uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	CourierID   		string      `gorm:"type:uuid;not null"`
	OriginState			string		`gorm:"size:100;not null"`
	OriginCity			string		`gorm:"size:100;not null"`
	DestinationState	string		`gorm:"size:100;not null"`
	DestinationCity		string		`gorm:"size:100;not null"`
	CreatedAt			time.Time	`gorm:"autoCreateTime"`
	UpdatedAt			time.Time	`gorm:"autoUpdateTime"`
}