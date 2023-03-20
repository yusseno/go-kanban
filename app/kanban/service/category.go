package service

import "go-kanban/app/kanban/repository"

type CategoryService interface {
	CreateCategory() error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (c *categoryService) CreateCategory() error {
	return nil
}
