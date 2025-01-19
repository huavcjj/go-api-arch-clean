package usecase

import (
	"go-api-arch-clean/adapter/gateway"
	"go-api-arch-clean/entity"
)

type (
	CategoryUseCase interface {
		GetOrCreate(category *entity.Category) (*entity.Category, error)
	}
)

type categoryUseCase struct {
	categoryRepository gateway.CategoryRepository
}

func NewCategoryUseCase(categoryRepository gateway.CategoryRepository) *categoryUseCase {
	return &categoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (a *categoryUseCase) GetOrCreate(category *entity.Category) (*entity.Category, error) {
	return a.categoryRepository.GetOrCreate(category)
}
