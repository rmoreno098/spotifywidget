package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"spotify-widget-v2/models"
	"time"
)

type JwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

type JwtClaims struct {
	jwt.RegisteredClaims
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
	Iat   int64  `json:"iat"`
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{
		secretKey:     secretKey,
		tokenDuration: time.Hour * 1, // fallback
	}
}

func (s *JwtService) GenerateToken(user *models.UserProfile) (string, error) {
	claims := JwtClaims{
		Sub:   user.ID,
		Name:  user.DisplayName,
		Email: user.Email,
		Exp:   time.Now().Add(s.tokenDuration).Unix(),
		Iat:   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JwtService) ParseToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	// Return better message
	return nil, errors.New("invalid token")
}
