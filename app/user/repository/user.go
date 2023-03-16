package repository

import (
	"errors"
	"go-kanban/app/user/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	UserLoginDB(user entity.User) (resUser entity.User, err error)
	UserRegisterDB(user entity.User) error
	UserGetByEmail(user entity.User) (email string, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) UserLoginDB(user entity.User) (resUser entity.User, err error) {
	res := r.db.Table("users").Where("email = ?", user.Email).Find(&resUser)
	if res.Error != nil {
		return entity.User{}, errors.New("select error")
	}
	if resUser.Email == "" {
		return entity.User{}, errors.New("user not found")
	}
	return resUser, nil
}

func (r *userRepository) UserRegisterDB(user entity.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) UserGetByEmail(user entity.User) (email string, err error) {
	res := r.db.Table("users").Select("email").Where("email = ?", user.Email).Scan(&email)
	if res.Error != nil {
		return "", errors.New("select error")
	}
	if email == "" {
		return "", errors.New("email empty")
	}
	return user.Email, nil
}
