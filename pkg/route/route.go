package route

import (
	"github.com/bluespada/timewise/internal/utils/types"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.All("/", func(ctx *fiber.Ctx) error {
		res := types.NewApiResponse()
		res.Message = "Timewise API"
		res.Data = map[string]interface{}{
			"version": "1.0.0",
		}
		return ctx.JSON(res)
	})
}
