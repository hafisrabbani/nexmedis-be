package service

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
	expiry time.Duration
}

func NewJWTService() *JWTService {
	return &JWTService{
		secret: []byte(os.Getenv("JWT_SECRET")),
		expiry: time.Minute * time.Duration(getJWTExpiry()),
	}
}

func getJWTExpiry() int {
	if v := os.Getenv("JWT_EXPIRED_MINUTES"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return 60
}

func (s *JWTService) Generate(clientID string) (string, error) {
	claims := jwt.MapClaims{
		"client_id": clientID,
		"exp":       time.Now().Add(s.expiry).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *JWTService) Validate(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
