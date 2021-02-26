package service

import (
	"github.com/gosimple/slug"
	"goshop/entity"
	"goshop/model"
	"goshop/repository"
)

type (
	CategoryService interface {
		CreateCategory(input entity.CategoryUserinput) (bool, error)
		UpdateCategory(input entity.CategoryUserinput) (bool, error)
		ListCategory() ([]model.Category, error)
		DeleteCategory(input entity.IdUserInput, delete_by int) (bool, error)
		DeletePermanent(input entity.IdUserInput) (bool, error)
		ListTrash() ([]model.Category, error)
		RestoreData(input entity.IdUserInput) (bool, error)
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
		Slug:      slug.Make(input.Name),
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
	category.Slug = slug.Make(input.Name)
	category.Image = &input.Image
	category.UpdatedBy = &input.UpdatedBy

	_, err = s.repository.UpdateCategory(category)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *category_service) ListCategory() ([]model.Category, error) {
	listCategory, err := s.repository.ListCategory()

	if err != nil {
		return listCategory, err
	}

	return listCategory, nil
}

func (s *category_service) DeleteCategory(input entity.IdUserInput, delete_by int) (bool, error) {
	_, err := s.repository.DeleteCategory(input.ID, delete_by)

	if err != nil {
		return false, err
	}

	return true, err
}

func (s *category_service) DeletePermanent(input entity.IdUserInput) (bool, error) {
	_, err := s.repository.DeletePermanent(input.ID)

	if err != nil {
		return false, err
	}

	return true, err
}

func (s *category_service) ListTrash() ([]model.Category, error) {
	data, err := s.repository.ListTrash()

	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *category_service) RestoreData(input entity.IdUserInput) (bool, error) {
	_, err := s.repository.RestoreData(input.ID)

	if err != nil {
		return false, err
	}

	return true, nil
}
