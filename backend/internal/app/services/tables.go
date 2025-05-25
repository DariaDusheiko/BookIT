package services

import (
	"time"
	
	"github.com/BookIT/backend/internal/app/models"
	"github.com/BookIT/backend/internal/app/repository"
)

type TableService interface {
	GetTablesWithOccupancy(start, end *time.Time) ([]TableWithOccupancy, error)
}

type tableService struct {
	tableRepo   repositories.TableRepository
	bookingRepo repositories.BookingRepository
}

type TableWithOccupancy struct {
	ID          uint       `json:"id"`
	X           int        `json:"x"`
	Y           int        `json:"y"`
	Angle       int        `json:"angle"`
	SeatsNumber int        `json:"seats_number"`
	Occupied    bool       `json:"occupied"`
	StartTime   *time.Time `json:"start_time,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
}

func NewTableService(
	tableRepo repositories.TableRepository, 
	bookingRepo repositories.BookingRepository,
) TableService {
	return &tableService{
		tableRepo:   tableRepo,
		bookingRepo: bookingRepo,
	}
}

func (s *tableService) GetTablesWithOccupancy(start, end *time.Time) ([]TableWithOccupancy, error) {
	tables, err := s.tableRepo.GetAllTables()
	if err != nil {
		return nil, err
	}

	var bookings []models.Booking
	if start != nil && end != nil {
		bookings, err = s.bookingRepo.GetBookingsInRange(*start, *end)
	} else if start != nil {
		bookings, err = s.bookingRepo.GetOngoingBookings(*start)
	} else {
		bookings = []models.Booking{}
	}

	if err != nil {
		return nil, err
	}

	occupiedMap := make(map[uint]bool)
	for _, b := range bookings {
		occupiedMap[b.TableID] = true
	}

	result := make([]TableWithOccupancy, 0, len(tables))
	for _, t := range tables {
		result = append(result, TableWithOccupancy{
			ID:          t.ID,
			X:           t.X,
			Y:           t.Y,
			Angle:       t.Angle,
			SeatsNumber: t.SeatsNumber,
			Occupied:    occupiedMap[t.ID],
			StartTime:   start,
			EndTime:     end,
		})
	}

	return result, nil
}