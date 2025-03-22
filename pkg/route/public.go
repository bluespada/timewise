package route

import (
	cauth "github.com/bluespada/timewise/pkg/controller/auth"
	"github.com/gofiber/fiber/v2"
)

func registerPublic(app fiber.Router) {
	// authentication route
	auth := app.Group("/auth")
	auth.Post("/signin", cauth.NewAuthController().PostAuthenticationSignIn)
	auth.Post("/signout", cauth.NewAuthController().PostAuthenticationSignOut)
	auth.Post("/signup", cauth.NewAuthController().PostAuthenticationSignUp)
	auth.Get("/session", cauth.NewAuthController().GetSession)
}
