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
		ListProductUser(userId int) ([]model.Product, error)
		UpdateProduct(product model.Product) error
		GetProductById(productId int) (model.Product, error)
		DeleteImageProduct(productID int) error
		DeleteProduct(productID int) error
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
	err := r.db.Where("user_id = ?", userId).Preload("ProductImage").Preload("Category").Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil

}

func (r *repository) GetProductById(productId int) (model.Product, error) {
	var product model.Product
	err := r.db.Preload("ProductImage").Where("id", productId).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) UpdateProduct(product model.Product) error {
	err := r.db.Save(&product).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteImageProduct(productID int) error {
	tx := r.db.Begin()
	err := tx.Where("product_id = ?", productID).Delete(&model.ProductImage{}).Error

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *repository) DeleteProduct(productID int) error {
	tx := r.db.Begin()
	errDeleteImage := r.DeleteImageProduct(productID)

	if errDeleteImage != nil {
		tx.Rollback()
		return errDeleteImage
	}

	err := tx.Where("id = ?", productID).Delete(model.Product{}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}
