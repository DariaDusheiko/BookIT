package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/BookIT/backend/config"
)

const (
	TokenTTLHours = 72 
)

var (
	jwtSecret string
)

func init() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	jwtSecret = cfg.App.SecretKey
}

func GenerateJWTToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(TokenTTLHours * time.Hour).Unix(),
	})

	return token.SignedString([]byte(jwtSecret)) 
}

