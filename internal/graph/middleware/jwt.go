package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

var (
	ErrNotFound = errors.New("Token not found")
	ErrNotValid = errors.New("Token not valid")
	ErrExpired  = errors.New("Token expired")
)

func ValidateJWT(c *fiber.Ctx) (jwt.MapClaims, error) {
	token := c.Get("Authorization")
	if token == "" {
		return nil, ErrNotFound
	}

	parts := strings.Split(token, "Bearer ")
	if len(parts) < 2 {
		return nil, ErrNotValid
	}

	tokenStr := strings.TrimSpace(parts[1])

	parsedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Debug Algoritma JWT
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrNotValid
		}

		return jwtSecret, nil
	})

	if err != nil {
		return nil, ErrNotValid
	}

	if !parsedToken.Valid {
		return nil, ErrNotValid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrExpired
	}

	return claims, nil
}

func init() {
	if os.Getenv("APP_JWT_SECRET") != "" {
		jwtSecret = []byte(os.Getenv("APP_JWT_SECRET"))
	} else {
		jwtSecret = []byte("4D3D621474572B7E35F615F5F9361")
	}
}
