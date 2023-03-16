package repository

import "gorm.io/gorm"

type UserRepository interface {
	UserLogin() error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) UserLogin() error {
	return nil
}
