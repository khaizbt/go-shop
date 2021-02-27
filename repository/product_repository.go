package repository

import (
	"goshop/config"
	"goshop/model"
)

type (
	ProductRepository interface {
		CreateProduct(product model.Product) (model.Product, error)
		StoreImage(image model.ProductImage) error
		ListProductAll() ([]model.Product, error)
	}
)

func NewProductCategory() *repository {
	return &repository{config.GetDB()}
}

func (r *repository) CreateProduct(product model.Product) (model.Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) StoreImage(image model.ProductImage) error {
	err := r.db.Create(&image).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListProductAll() ([]model.Product, error) {
	var product []model.Product
	err := r.db.Preload("ProductImage").Preload("Category").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) ListProductUser(userId int) ([]model.Product, error) {
	var product []model.Product
	err := r.db.Where("user_id = ?", userId).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil

}
