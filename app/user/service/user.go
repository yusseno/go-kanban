package service

import (
	"errors"
	"go-kanban/app/user/entity"
	"go-kanban/app/user/repository"
	"go-kanban/security"
)

type UserService interface {
	UserLogin(entity.User) (entity.User, error)
	UserRegister(entity.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) UserLogin(user entity.User) (entity.User, error) {
	resUser, err := u.userRepository.UserLoginDB(user)
	if err != nil {
		return entity.User{}, err
	}
	validator := security.CheckPasswordHash(user.Password, resUser.Password)
	if !validator {
		return entity.User{}, errors.New("password not match")
	}
	return resUser, nil
}

func (u *userService) UserRegister(user entity.User) error {
	_, err := u.userRepository.UserGetByEmail(user)
	if err == nil {
		return errors.New("user already exists")
	}

	hash, err := security.HashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = hash
	err = u.userRepository.UserRegisterDB(user)
	if err != nil {
		return errors.New("create to databases failed")
	}
	return nil
}
