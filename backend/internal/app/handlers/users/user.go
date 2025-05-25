package users

import (
	"net/http"
	"github.com/BookIT/backend/internal/app/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
	schemas     *AuthSchemas
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		schemas:     &AuthSchemas{},
	}
}

// Authenticate godoc
// @Summary User authentication/registration
// @Description Authenticates existing user or creates new user if phone number doesn't exist. Returns JWT token for authorized requests.
// @Tags auth
// @Accept json
// @Produce json
// @Param input body AuthRequest true "User credentials"
// @Success 200 {object} AuthResponse "Successfully authenticated/registered"
// @Failure 400 {object} ErrorResponse "Invalid request format or validation error"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth [post]
func (h *UserHandler) Authenticate(c *gin.Context) {
	req, err := h.schemas.ValidateAuthRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.schemas.NewErrorResponse(err))
		return
	}

	token, err := h.userService.AuthenticateOrRegister(req.Username, req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, h.schemas.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, AuthResponse{Token: token})
}