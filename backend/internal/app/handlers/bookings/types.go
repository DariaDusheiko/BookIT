package bookings

// все еще ждем прибавления....

import "time"

type CreateBookingRequest struct {
	TableID   uint       `json:"table_id" binding:"required"`
	StartTime time.Time  `json:"start_time" binding:"required"`
	EndTime   *time.Time `json:"end_time"`
}

type BookingResponse struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	TableID   uint       `json:"table_id"`
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type DeleteBookingRequest struct {
	BookingID uint `json:"booking_id" binding:"required"`
}

type DeleteBookingResponse struct {
	Message string `json:"message"`
}
