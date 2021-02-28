package service

import (
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"goshop/entity"
	"goshop/model"
	"goshop/repository"
	"log"
	"os"
)

type (
	ProductService interface {
		CreateProduct(input entity.ProductInput) error
		ListProduct() ([]model.Product, error)
		ListProductUser(userID int) ([]model.Product, error)
		UpdateProduct(url entity.IdUserInput, input entity.ProductInput) error
		DeleteProduct(input entity.IdUserInput) error
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

func (s *productService) ListProductUser(userID int) ([]model.Product, error) {
	listProduct, err := s.repository.ListProductUser(userID)

	if err != nil {
		return nil, err
	}

	return listProduct, nil
}

func (s *productService) UpdateProduct(url entity.IdUserInput, input entity.ProductInput) error {
	getProduct, err := s.repository.GetProductById(url.ID)

	if err != nil {
		return err
	}

	if url.User.ID != getProduct.UserID {
		return errors.New("You dont have access to this product")
	}

	getProduct.Title = input.Title
	getProduct.CategoryId = input.CategoryID
	getProduct.Slug = slug.Make(input.Title)
	getProduct.PurchasePrice = input.PurchasePrice
	getProduct.SellingPrice = input.SellingPrice
	getProduct.Description = input.Description
	getProduct.InstagramAccount = input.InstagramAccount
	getProduct.Phone = input.Phone
	getProduct.Address = input.Address

	err = s.repository.UpdateProduct(getProduct)

	if err != nil {
		return err
	}

	if input.Image != nil {
		err := s.repository.DeleteImageProduct(getProduct.ID)
		if err != nil {
			return err
		}
		for i, image := range input.Image {
			e := os.Remove(getProduct.ProductImage[i].ImageName)
			if e == nil {
				log.Fatal("Failed delete image")
			}

			var productImage model.ProductImage
			if image.IsPrimary {
				productImage.IsPrimary = true
			}

			productImage.ProductID = getProduct.ID
			productImage.ImageName = image.ImageName

			err := s.repository.StoreImage(productImage)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *productService) DeleteProduct(input entity.IdUserInput) error {
	//Delete File Product from storage
	product, _ := s.repository.GetProductById(input.ID)
	if product.ID == 0 {
		return errors.New("Data Not Found")
	}

	if input.User.ID != product.UserID {
		return errors.New("You dont have access to this product")
	}
	for _, image := range product.ProductImage {
		e := os.Remove(image.ImageName)

		if e != nil {
			panic("Error delete File image")
		}
	}

	err := s.repository.DeleteProduct(input.ID)

	if err != nil {
		return err
	}

	return nil
}
