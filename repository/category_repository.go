package repository

import (
	"goshop/config"
	"goshop/model"
)

type (
	CategoryRepository interface {
		CreateCategory(category model.Category) (model.Category, error)
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
