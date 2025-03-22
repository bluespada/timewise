package main

import (
	"github.com/bluespada/timewise/internal/utils/helper"
	"github.com/bluespada/timewise/pkg/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// parsing sysmtem env to here

	app := fiber.New(fiber.Config{})
	viteAssets := helper.GetViteMetadata("./web/dist/.vite/manifest.json")
	if len(viteAssets) > 0 {
		println("vite assets not found")
	}
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())
	// register the route
	route.RegisterRoutes(app)
	app.Listen(":8000")
}
