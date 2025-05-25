package bookings

// ждем прибавления....

import (
	"github.com/BookIT/backend/internal/app/models"
	"github.com/gin-gonic/gin"
)

type BookingSchemas struct{}

func (s *BookingSchemas) ValidateCreateRequest(c *gin.Context) (*CreateBookingRequest, error) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (s *BookingSchemas) NewBookingResponse(booking *models.Booking) *BookingResponse {
	return &BookingResponse{
		ID:        booking.ID,
		UserID:    booking.UserID,
		TableID:   booking.TableID,
		StartTime: booking.StartTime,
		EndTime:   booking.EndTime,
	}
}

func (s *BookingSchemas) NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Error: message}
}