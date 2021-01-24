package service

import (
	"goshop/entity"
	"goshop/model"
	"goshop/repository"
)

type (
	CategoryService interface {
	}

	category_service struct {
		repository repository.CategoryRepository
	}
)

func NewCategoryService(repository repository.CategoryRepository) *category_service {
	return &category_service{repository}
}

func (s *category_service) CreateCategory(input entity.CategoryUserinput) (bool, error) {
	category := model.Category{
		Name:      input.Name,
		Slug:      &input.Slug,
		Image:     &input.Image,
		CreatedBy: input.CreatedBy,
	}

	_, err := s.repository.CreateCategory(category)

	if err != nil {
		return false, err
	}

	return true, err
}
