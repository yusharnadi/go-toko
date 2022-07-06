package service

import (
	"github.com/yusharnadi/go-toko/entity"
	"github.com/yusharnadi/go-toko/repository"
)

type authService struct {
	authRepository repository.AuthRepository
}

type AuthService interface {
	FindByEmail(email string, user *entity.User) (*entity.User, error)
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository}
}

func (s *authService) FindByEmail(email string, user *entity.User) (*entity.User, error) {
	user, err := s.authRepository.FindByEmail(email, user)

	return user, err
}
