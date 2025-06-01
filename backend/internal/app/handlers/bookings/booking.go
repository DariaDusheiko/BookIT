package bookings

// 1 этап в нашей табличке из лекций!!!!!
// Валидация данных и получения запроса

// тут необходимо добавить 2 функции, которые относятся к
// 1 - удалению записи
// 2 - получении информации о бронированиях пользователя

// связать валидацию данных на данном этапе на преобразование
// к необходмому формат json для отправки на фронт
//
// (добавить соответствующие схемы и типы для твоих функций) для запроса и ответа

// bookingGroup.DELETE("/", bookingHandler.DeleteBooking)
// bookingGroup.POST("info/", bookingHandler.InfoBooking)

// такие названия функций лежат в backend/cmd/server/main
// если названия будут отличаться, то необходимо поменять
// и привязать роутер корректно

// учти, что все апи долдны быть защищены с помощью токена
// тебе необходимо только вставлять

// userID, exists := c.Get("userID")
//     if !exists {
//         c.JSON(401, h.schemas.NewErrorResponse("authentication required"))
//         return
//     }

// в начало каждой функкции чтобы проверять нормально ли все

// отсюда также ты получешь информацию об ИД пользователя
// которое тебе вохможно придется использовать в sql запросах

// не забудь писать тут информацию для сваггера

import (
	"github.com/BookIT/backend/internal/app/services"
	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	service services.BookingService
	schemas *BookingSchemas
}

func NewBookingHandler(service services.BookingService) *BookingHandler {
	return &BookingHandler{
		service: service,
		schemas: &BookingSchemas{},
	}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new table booking for authenticated user
// @Tags bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param X-Auth-Token header string true "JWT Token"
// @Param request body CreateBookingRequest true "Booking details"
// @Success 200 {object} BookingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, h.schemas.NewErrorResponse("authentication required"))
		return
	}

	req, err := h.schemas.ValidateCreateRequest(c)
	if err != nil {
		c.JSON(400, h.schemas.NewErrorResponse(err.Error()))
		return
	}

	booking, err := h.service.Create(
		userID.(uint),
		req.TableID,
		req.StartTime,
		req.EndTime,
	)

	if err != nil {
		c.JSON(500, h.schemas.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(200, h.schemas.NewBookingResponse(booking))
}

// DeleteBooking godoc
// @Summary Delete a booking
// @Description Delete user's booking by ID
// @Tags bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param X-Auth-Token header string true "JWT Token"
// @Param request body DeleteBookingRequest true "Booking ID"
// @Success 200 {object} DeleteBookingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings [delete]
func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	// Авторизация пользователя
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, h.schemas.NewErrorResponse("authentication required"))
		return
	}

	// Валидация запроса
	req, err := h.schemas.ValidateDeleteRequest(c)
	if err != nil {
		c.JSON(400, h.schemas.NewErrorResponse(err.Error()))
		return
	}

	// Вызов запроса
	err = h.service.Delete(
		userID.(uint),
		req.BookingID,
	)

	// Ответ
	if err != nil {
		c.JSON(500, h.schemas.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(200, h.schemas.NewDeleteResponse())
}

// GetUserBookings godoc
// @Summary Get user bookings
// @Description Get all bookings for authenticated user
// @Tags bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param X-Auth-Token header string true "JWT Token"
// @Success 200 {object} UserBookingsResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /bookings/info [get]
func (h *BookingHandler) GetUserBookings(c *gin.Context) {
	// Авторизация пользователя
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, h.schemas.NewErrorResponse("authentication required"))
		return
	}

	// Получение бронирований через сервис
	bookings, err := h.service.GetUserBookings(userID.(uint))
	if err != nil {
		c.JSON(500, h.schemas.NewErrorResponse(err.Error()))
		return
	}

	// Формирование ответа
	c.JSON(200, h.schemas.NewUserBookingsResponse(bookings))
}
