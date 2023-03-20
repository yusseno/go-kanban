package repository

import "gorm.io/gorm"

type CategoryRepository interface {
	CreateCategoryDB() error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategoryDB() error {
	return nil
}
