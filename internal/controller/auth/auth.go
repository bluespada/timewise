// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains authentication controller
package auth

import (
	"os"
	"time"

	"github.com/bluespada/timewise/internal/repositories"
	"github.com/bluespada/timewise/internal/utils/database"
	"github.com/bluespada/timewise/internal/utils/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

type SignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// HandleSignIn handle signin request from client. It will generate JWT token that
// will be used for authentication in private route. The token will be expired in
// 14 days.
func HandleSignIn(c *fiber.Ctx) error {

	authRepo := repositories.NewAuthRepositories(database.Db)

	res := types.NewApiResponse()

	var params SignInParams
	if err := c.BodyParser(&params); err != nil {
		res.Error = true
		res.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	auth, err := authRepo.FindByEmail(params.Email)
	if err != nil {
		res.Error = true
		res.Message = "Account Not found."
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 336).Unix(),
		"iat":  time.Now().Unix(),
		"user": auth.ID,
		"sub":  "user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	res.Data = map[string]any{
		"token": t,
		"auth":  auth,
	}
	return c.JSON(res)
}

func init() {
	if os.Getenv("APP_JWT_SECRET") != "" {
		jwtSecret = os.Getenv("APP_JWT_SECRET")
	} else {
		jwtSecret = "4D3D621474572B7E35F615F5F9361"
	}
}
