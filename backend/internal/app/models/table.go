package models

type Table struct {
	ID          uint `gorm:"primaryKey"`
	X           int  `gorm:"not null"`
	Y           int  `gorm:"not null"`
	Angle       int  `gorm:"not null"`
	SeatsNumber int  `gorm:"not null"`
}
