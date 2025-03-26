package repositories

import (
	"github.com/bluespada/timewise/internal/model"
	"gorm.io/gorm"
)

type AuthRepositories struct {
	db *gorm.DB
}

func NewAuthRepositories(db *gorm.DB) *AuthRepositories {
	return &AuthRepositories{db}
}

func (r *AuthRepositories) FindByEmail(email string) (model.ModelAuth, error) {
	auth := model.ModelAuth{}
	dbr := r.db.Find(&auth, "email = ?", email)
	if dbr.RowsAffected == 0 {
		return auth, ErrNotFound
	}
	return auth, nil
}

func (r *AuthRepositories) FindByPhone(phone string) (model.ModelAuth, error) {
	auth := model.ModelAuth{}
	dbr := r.db.Find(&auth, "phone = ?", phone)
	if dbr.RowsAffected == 0 {
		return auth, ErrNotFound
	}
	return auth, nil
}
