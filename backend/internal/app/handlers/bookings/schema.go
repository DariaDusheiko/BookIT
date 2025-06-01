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

func (s *BookingSchemas) ValidateDeleteRequest(c *gin.Context) (*DeleteBookingRequest, error) {
	var req DeleteBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (s *BookingSchemas) NewDeleteResponse() *DeleteBookingResponse {
	return &DeleteBookingResponse{
		Message: "Booking deleted successfully",
	}
}

func (s *BookingSchemas) NewUserBookingsResponse(bookings []models.Booking) *UserBookingsResponse {
	result := make([]BookingResponse, 0, len(bookings))
	for _, b := range bookings {
		result = append(result, *s.NewBookingResponse(&b))
	}
	return &UserBookingsResponse{Bookings: result}
}
