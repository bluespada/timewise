package services

type AuthServices struct{}

func NewAuthService() *AuthServices {
	return &AuthServices{}
}

func (as *AuthServices) CreateUser() {
	// TODO: Transaction to create a auth and the users.
}
