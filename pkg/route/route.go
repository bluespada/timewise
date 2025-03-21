package route

import (
	"github.com/bluespada/timewise/internal/utils/types"
	cauth "github.com/bluespada/timewise/pkg/controller/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	api := app.Group("/api")
	public := api.Group("/public")
	auth := public.Group("/auth")

	// Authentication routers
	auth.Post("/signin", cauth.NewAuthController().PostAuthenticationSignIn)
	auth.Post("/signout", cauth.NewAuthController().PostAuthenticationSignOut)
	auth.Post("/signup", cauth.NewAuthController().PostAuthenticationSignUp)
	auth.Get("/session", cauth.NewAuthController().GetSession)

	api.All("/", func(ctx *fiber.Ctx) error {
		res := types.NewApiResponse()
		res.Message = "Timewise API"
		res.Data = map[string]interface{}{
			"version": "1.0.0",
		}
		return ctx.JSON(res)
	})
}
