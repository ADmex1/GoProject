package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/ADMex1/GoProject/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Generate Token Generate Refresh Token
func GenerateToken(userID int64, role string, email string, publicID uuid.UUID) (string, error) {
	secret := strings.TrimSpace(config.AppConfig.JWTSecret)
	if secret == "" {
		return "", fmt.Errorf("JWT secret is empty")
	}

	duration, err := time.ParseDuration(config.AppConfig.JWTExpire)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id":  userID,
		"role":     role,
		"pub_id":   publicID,
		"email":    email,
		"exp_time": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
func RefreshToken(userID int64) (string, error) {
	secret := strings.TrimSpace(config.AppConfig.JWTSecret)
	if secret == "" {
		return "", fmt.Errorf("JWT secret is empty")
	}

	duration, err := time.ParseDuration(config.AppConfig.JWTRefreshToken)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id":  userID,
		"exp_time": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
