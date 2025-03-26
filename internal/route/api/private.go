// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains routing for private api route
package api

import (
	"os"

	"github.com/bluespada/timewise/internal/utils/types"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

// InitPrivateRoute initializes the private API routes with JWT authentication middleware.
// The middleware checks the validity of the JWT token in the request headers.
// If the token is invalid, an unauthorized error response is returned.
// On successful authentication, routes under this group can be accessed.
func InitPrivateRoute(app fiber.Router) {

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(jwtSecret),
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			res := types.NewApiResponse()
			res.Error = true
			res.Message = err.Error()
			return c.Status(fiber.StatusUnauthorized).JSON(res)
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		res := types.NewApiResponse()
		res.Message = "You are authenticated"
		res.Data = claims
		return c.JSON(res)
	})

}

func init() {
	if os.Getenv("APP_JWT_SECRET") != "" {
		jwtSecret = os.Getenv("APP_JWT_SECRET")
	} else {
		jwtSecret = "4D3D621474572B7E35F615F5F9361"
	}
}
