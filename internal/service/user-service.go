package service

import (
	"webfolio-backend/internal/domain"
	"webfolio-backend/internal/repository"
)

type UserService interface {
	CreateUser(user *domain.User) (*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) CreateUser(user *domain.User) (*domain.User, error) {
	return s.userRepo.Create(user)
}
