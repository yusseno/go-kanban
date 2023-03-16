package service

import "go-kanban/app/repository"

type UserService interface {
	UserLogin()
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) UserLogin() {}
