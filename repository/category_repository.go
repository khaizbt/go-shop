package repository

import (
	"goshop/config"
	"goshop/model"
)

type (
	CategoryRepository interface {
		CreateCategory(category model.Category) (model.Category, error)
		FindCategoryByID(categoryID int) (model.Category, error)
		UpdateCategory(category model.Category) (bool, error)
	}
)

func NewCategoryRepository() *repository {
	return &repository{config.GetDB()}
}

func (r *repository) CreateCategory(category model.Category) (model.Category, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindCategoryByID(categoryID int) (model.Category, error) {
	var category model.Category
	err := r.db.Where("id = ?", categoryID).First(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) UpdateCategory(category model.Category) (bool, error) {
	err := r.db.Save(&category).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

//TODO List Category
//
//func (r *repository) ListCategory() ([]model.Category, error) {
//	var category mo
//}
