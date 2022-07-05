package service

import (
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Insert(product *entity.User) error
	GetAll() (*[]entity.User, error)
	FindId(id int) (entity.User, error)
	Update(product *entity.User, id int) error
	Delete(id int) error
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Insert(user *entity.User) error {

	err := s.userRepository.Insert(user)

	return err
}

func (s *userService) GetAll() (*[]entity.User, error) {
	users, err := s.userRepository.GetAll()

	return users, err
}

func (s *userService) FindId(id int) (entity.User, error) {
	data, err := s.userRepository.FindId(id)

	return data, err
}

func (s *userService) Update(user *entity.User, id int) error {
	err := s.userRepository.Update(user, id)

	return err
}

func (s *userService) Delete(id int) error {
	return s.userRepository.Delete(id)
}
