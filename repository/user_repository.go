package repository

import (
	"github.com/yusharnadi/go-toko/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *entity.User) error
	GetAll() (*[]entity.User, error)
	FindId(id int) (entity.User, error)
	Update(user *entity.User, id int) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Insert(user *entity.User) error {
	err := r.db.Create(&user)

	return err.Error
}

func (r *userRepository) GetAll() (*[]entity.User, error) {
	var user *[]entity.User
	err := r.db.Find(&user)

	return user, err.Error
}

func (r *userRepository) FindId(id int) (entity.User, error) {
	var user entity.User
	err := r.db.Find(&user, id)

	return user, err.Error
}

func (r *userRepository) Update(user *entity.User, id int) error {
	var User *entity.User
	err := r.db.Model(&User).Where("id", id).Updates(user)

	return err.Error
}

func (r *userRepository) Delete(id int) error {
	var User *entity.User
	err := r.db.Delete(&User, id)
	return err.Error
}
