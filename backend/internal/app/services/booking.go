package services

// 2 этап в нашей табличке из лекций!!!!!
// Описание бизнес логики, что происходит

// описать какие запросы и с какой последовательности выполнить к бд
// что из них взять и что лепить

// тут также емть интерфейс но я не стала их в отдельный файл класть уже
// хотя надо бы наверное

import (
	"errors"
	"time"
	
	"github.com/BookIT/backend/internal/app/models"
	"github.com/BookIT/backend/internal/app/repository"
)

type BookingService interface {
	Create(userID, tableID uint, start time.Time, end *time.Time) (*models.Booking, error)
	IsTableAvailable(tableID uint, start time.Time, end *time.Time) (bool, error)
}

type bookingService struct {
	repo      repositories.BookingRepository
	tableRepo repositories.TableRepository
}

func NewBookingService(repo repositories.BookingRepository, tableRepo repositories.TableRepository) BookingService {
	return &bookingService{
		repo:      repo,
		tableRepo: tableRepo,
	}
}

func (s *bookingService) Create(userID, tableID uint, start time.Time, end *time.Time) (*models.Booking, error) {
	available, err := s.IsTableAvailable(tableID, start, end)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, errors.New("table is not available for the selected time slot")
	}

	booking := &models.Booking{
		UserID:    userID,
		TableID:   tableID,
		StartTime: start,
		EndTime:   end,
	}

	if err := s.repo.CreateBooking(booking); err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *bookingService) IsTableAvailable(tableID uint, start time.Time, end *time.Time) (bool, error) {
	var bookings []models.Booking
	var err error

	if end != nil {
		bookings, err = s.repo.GetBookingsForTable(tableID, start, *end)
	} else {
		bookings, err = s.repo.GetOngoingBookingsForTable(tableID, start)
	}

	if err != nil {
		return false, err
	}

	return len(bookings) == 0, nil
}