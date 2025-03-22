package route

import (
	"github.com/bluespada/timewise/internal/utils/types"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	api := app.Group("/api")

	// asignning public and private routing
	registerPublic(api.Group("/public"))
	registerPrivate(api.Group("/private"))

	// handle index routing
	api.All("/", func(ctx *fiber.Ctx) error {
		res := types.NewApiResponse()
		res.Message = "Timewise API"
		res.Data = map[string]interface{}{
			"version": "1.0.0",
		}
		return ctx.JSON(res)
	})

	// set not found pages
	api.All("*", func(c *fiber.Ctx) error {
		res := types.NewApiResponse()
		res.Error = true
		res.Message = "Not Found."
		return c.Status(fiber.StatusNotFound).JSON(res)
	})
}
