package service

import (
	"goshop/entity"
	"goshop/model"
	"goshop/repository"
)

type (
	CategoryService interface {
		CreateCategory(input entity.CategoryUserinput) (bool, error)
		UpdateCategory(input entity.CategoryUserinput) (bool, error)
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

func (s *category_service) UpdateCategory(input entity.CategoryUserinput) (bool, error) {
	category, err := s.repository.FindCategoryByID(input.ID)

	if err != nil {
		return false, err
	}

	category.Name = input.Name
	category.Slug = &input.Slug
	category.Image = &input.Image
	category.UpdatedBy = &input.UpdatedBy

	_, err = s.repository.UpdateCategory(category)

	if err != nil {
		return false, err
	}

	return true, nil

}
