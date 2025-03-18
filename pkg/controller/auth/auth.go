package auth

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) PostAuthenticationSignIn() error {
	return nil
}

func (ac *AuthController) PostAuthenticationSignUp() error {
	return nil
}

func (ac *AuthController) PostAuthenticationSignOut() error {
	return nil
}
