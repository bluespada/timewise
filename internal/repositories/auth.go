package repositories

import (
	"github.com/bluespada/timewise/internal/model"
	"gorm.io/gorm"
)

type AuthRepositories struct {
	db *gorm.DB
}

// NewAuthRepositories will create an instance of AuthRepositories. This is the
// primary entry-point for using the auth repository.
func NewAuthRepositories(db *gorm.DB) *AuthRepositories {
	return &AuthRepositories{db}
}

// FindByEmail will find user by email. If user not found, it will return ErrNotFound error.
func (r *AuthRepositories) FindByEmail(email string) (model.ModelAuth, error) {
	auth := model.ModelAuth{}
	dbr := r.db.Find(&auth, "email = ?", email)
	if dbr.RowsAffected == 0 {
		return auth, ErrNotFound
	}
	return auth, nil
}

// FindByPhone will find user by phone. If user not found, it will return ErrNotFound error.
func (r *AuthRepositories) FindByPhone(phone string) (model.ModelAuth, error) {
	auth := model.ModelAuth{}
	dbr := r.db.Find(&auth, "phone = ?", phone)
	if dbr.RowsAffected == 0 {
		return auth, ErrNotFound
	}
	return auth, nil
}

// All will return all user. If no user found, it will return ErrNotFound error.
func (r *AuthRepositories) All() ([]model.ModelAuth, error) {
	var auths []model.ModelAuth
	dbr := r.db.Find(&auths)
	if dbr.RowsAffected == 0 {
		return auths, ErrNotFound
	}
	return auths, nil
}
