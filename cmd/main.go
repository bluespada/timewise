package main

import (
	"github.com/bluespada/timewise/pkg/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{})
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())
	// register the route
	route.RegisterRoutes(app)
	app.Listen(":8000")
}
