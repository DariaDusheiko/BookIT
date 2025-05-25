package services

import (
	"errors"
	
	"github.com/BookIT/backend/internal/app/models"
	"github.com/BookIT/backend/internal/app/repository"
	"github.com/BookIT/backend/internal/pkg/utils"
)

type UserService interface {
	AuthenticateOrRegister(username, phoneNumber string) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) AuthenticateOrRegister(username, phoneNumber string) (string, error) {
	user, err := s.userRepo.FindByPhoneNumber(phoneNumber)
	if err == nil {
		return utils.GenerateJWTToken(user.ID)
	}

	newUser := &models.User{
		Username:    username,
		PhoneNumber: phoneNumber,
	}

	if err := s.userRepo.Create(newUser); err != nil {
		return "", errors.New("failed to create user")
	}

	return utils.GenerateJWTToken(newUser.ID)
}