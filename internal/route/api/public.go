// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains routing for public api route
package api

import (
	auth_controller "github.com/bluespada/timewise/internal/controller/auth"
	"github.com/gofiber/fiber/v2"
)

// InitPublicApiRoute is a function that will initialize all public api route.
func InitPublicApiRoute(app fiber.Router) {

	// auth route
	auth_route := app.Group("/auth")
	auth_route.Post("/signin", auth_controller.HandleSignIn)
}
