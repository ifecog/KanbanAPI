package models

import "time"

type PackageSize string

const (
	Small  PackageSize = "SMALL"
	Medium PackageSize = "MEDIUM"
	Large  PackageSize = "LARGE"
)

type Package struct {
	ID          		uint        `gorm:"primaryKey"`
	UUID        		string      `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	CourierID   		string      `gorm:"type:uuid;not null"`
	PackageSize 		PackageSize `gorm:"type:enum('SMALL', 'MEDIUM', 'LARGE')"`
	OriginState			string		`gorm:"size:100;not null"`
	OriginCity			string		`gorm:"size:100;not null"`
	DestinationState	string		`gorm:"size:100;not null"`
	DestinationCity		string		`gorm:"size:100;not null"`
	CreatedAt			time.Time	`gorm:"autoCreateTime"`
	UpdatedAt			time.Time	`gorm:"autoUpdateTime"`
}