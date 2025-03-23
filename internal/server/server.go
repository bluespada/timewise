package server

import (
	"github.com/bluespada/timewise/internal/route"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func RunApp() {
	// load dotenv
	godotenv.Load()

	// initialize gofiber
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	// global middleware
	app.Use(logger.New(logger.ConfigDefault))

	// initialize route
	route.InitRoute(app)

	// http listen
	app.Listen(":8000")

}
