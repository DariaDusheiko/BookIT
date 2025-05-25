package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type AuthSchemas struct{}

func (s *AuthSchemas) ValidateAuthRequest(c *gin.Context) (*AuthRequest, error) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	
	return &req, nil
}

func (s *AuthSchemas) NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}