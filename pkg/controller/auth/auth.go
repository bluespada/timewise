package auth

import "github.com/gofiber/fiber/v2"

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) PostAuthenticationSignIn(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) PostAuthenticationSignUp(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) PostAuthenticationSignOut(c *fiber.Ctx) error {
	return nil
}

func (ac *AuthController) GetSession(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    nil,
	})
}
