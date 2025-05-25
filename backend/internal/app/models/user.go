package models

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"not null"`
    PhoneNumber string `gorm:"unique;not null"`
}