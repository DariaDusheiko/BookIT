package models
import "time"

// не меняй тут ничего, а то мои апи тоже связаны с данной моделькой 
// тут ты описал уже типы и названия переменных для работы с таблицами в коде

type Booking struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index"`
	TableID   uint      `gorm:"index"`
	StartTime time.Time `gorm:"not null"`
	EndTime   *time.Time
}