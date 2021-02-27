package service

import (
	"github.com/gosimple/slug"
	"goshop/entity"
	"goshop/model"
	"goshop/repository"
)

type (
	ProductService interface {
		CreateProduct(input entity.ProductInput) error
		ListProduct() ([]model.Product, error)
	}

	productService struct {
		repository repository.ProductRepository
	}
)

func NewProductService(repository repository.ProductRepository) *productService {
	return &productService{repository: repository}
}

func (s *productService) CreateProduct(input entity.ProductInput) error {
	product := model.Product{
		Title:            input.Title,
		Description:      input.Description,
		Address:          input.Address,
		Phone:            input.Phone,
		InstagramAccount: input.InstagramAccount,
		PurchasePrice:    input.PurchasePrice,
		SellingPrice:     input.SellingPrice,
		CategoryId:       input.CategoryID,
		UserID:           input.UserID,
	}

	product.Slug = slug.Make(input.Title)

	product, err := s.repository.CreateProduct(product)

	if err != nil {
		return err
	}

	if input.Image != nil {
		for _, image := range input.Image {
			var productImage model.ProductImage
			if !image.IsPrimary {
				productImage.IsPrimary = false
			}

			productImage.ProductID = product.ID
			productImage.ImageName = image.ImageName

			err := s.repository.StoreImage(productImage)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *productService) ListProduct() ([]model.Product, error) {
	listProduct, err := s.repository.ListProductAll()

	if err != nil {
		return nil, err
	}

	return listProduct, err

}

//TODO List Product by User
