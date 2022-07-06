package repository

import (
	"github.com/yusharnadi/go-toko/entity"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	FindByEmail(email string, user *entity.User) (*entity.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) FindByEmail(email string, user *entity.User) (*entity.User, error) {

	err := r.db.Where("email", email).Find(&user).Error

	return user, err
}
