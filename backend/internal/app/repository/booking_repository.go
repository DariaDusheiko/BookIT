package repositories

// 3 этап в нашей табличке из лекций!!!!!

// а тут с тобой мы работает с таблицами sql через код
// необходимо добавить для твоих функций

// ВАЖНО! заметить, что для корректной работы необходимо испольовать интерфейсы
//
// используй backend/internal/app/repository/interfaces.go
//
// которые помогут определить, что пришло в функцию и что должно уйти
// так что добавил функцию тут - добавил там описания данных

import (
	"errors"
	"time"

	"github.com/BookIT/backend/internal/app/models"
	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetBookingsForTable(tableID uint, start, end time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("table_id = ?", tableID).
		Where("(start_time <= ? AND (end_time IS NULL OR end_time >= ?)) OR "+
			"(start_time <= ? AND (end_time IS NULL OR end_time >= ?))",
			end, end,
			start, start).
		Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetOngoingBookingsForTable(tableID uint, from time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("table_id = ?", tableID).
		Where("start_time <= ? AND end_time IS NULL", from).
		Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetBookingsInRange(start, end time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where(
		"(start_time <= ? AND (end_time IS NULL OR end_time >= ?)) OR "+
			"(start_time <= ? AND (end_time IS NULL OR end_time >= ?))",
		end, end,
		start, start).
		Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetOngoingBookings(from time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("start_time <= ? AND end_time IS NULL", from).
		Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) GetBookingByID(bookingID uint) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.Where("id = ?", bookingID).First(&booking).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Бронь не найдена
	}
	return &booking, err
}

func (r *bookingRepository) DeleteBooking(bookingID uint) error {
	return r.db.Where("id = ?", bookingID).Delete(&models.Booking{}).Error
}
