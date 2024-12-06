package service

import (
	"webfolio-backend/domain"
	"webfolio-backend/repository"
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
