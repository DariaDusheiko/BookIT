package repositories

import (
	"gorm.io/gorm"
	"github.com/BookIT/backend/internal/app/models"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}