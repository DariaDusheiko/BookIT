package repositories

// ждем прибавления описания типов переменных пришедших общаться с бд....
// именно после изменений файла выше надо и тут все добавить данные!

import (
	"time"

	"github.com/BookIT/backend/internal/app/models"
)

type UserRepository interface {
	FindByPhoneNumber(phoneNumber string) (*models.User, error)
	Create(user *models.User) error
}

type BookingRepository interface {
	CreateBooking(booking *models.Booking) error
	GetBookingsForTable(tableID uint, start, end time.Time) ([]models.Booking, error)
	GetOngoingBookingsForTable(tableID uint, from time.Time) ([]models.Booking, error)
	GetBookingsInRange(start, end time.Time) ([]models.Booking, error)
	GetOngoingBookings(from time.Time) ([]models.Booking, error)
	GetBookingByID(bookingID uint) (*models.Booking, error)
	DeleteBooking(bookingID uint) error
}

type TableRepository interface {
	CreateTable(table *models.Table) error
	DeleteTableByID(id uint) error
	GetAllTables() ([]models.Table, error)
	GetTableByID(id uint) (*models.Table, error)
	CreateTables(tables []models.Table) ([]uint, error)
}
