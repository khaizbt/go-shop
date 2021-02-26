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
		ListCategory() ([]model.Category, error)
		DeleteCategory(categoryID int, delete_by int) (bool, error)
		ListTrash() ([]model.Category, error)
		DeletePermanent(categoryID int) (bool, error)
		RestoreData(categoryID int) (bool, error)
	}
)

func NewCategoryRepository() *repository {
	return &repository{config.GetDB()}
}

func (r *repository) CreateCategory(category model.Category) (model.Category, error) {
	tx := r.db.Begin()
	err := tx.Create(&category).Error
	if err != nil {
		tx.Rollback()
		return category, err
	}

	tx.Commit()
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
	tx := r.db.Begin()
	err := tx.Save(&category).Error

	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}

//TODO List Category
//
func (r *repository) ListCategory() ([]model.Category, error) {
	var category []model.Category

	err := r.db.Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) DeleteCategory(categoryID int, delete_by int) (bool, error) {
	tx := r.db.Begin()
	data := tx.Where("id = ?", categoryID)
	data.Model(&model.Category{}).Update("deleted_by", &delete_by)
	err := data.Delete(&model.Category{}).Error

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}

func (r *repository) ListTrash() ([]model.Category, error) {
	var category []model.Category
	err := r.db.Unscoped().Where("deleted_at != ?", nil).Find(&category).Error

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *repository) DeletePermanent(categoryID int) (bool, error) {
	tx := r.db.Begin()
	err := tx.Unscoped().Where("id", 2).Delete(&model.Category{}).Error

	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}

func (r *repository) RestoreData(categoryID int) (bool, error) {
	err := r.db.Unscoped().Model(&model.Category{}).Where("id", categoryID).Update("deleted_at", nil).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
