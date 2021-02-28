package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"goshop/entity"
	"goshop/helper"
	"goshop/model"
	"goshop/service"
	"net/http"
	"strconv"
)

type productController struct {
	service service.ProductService
}

func ProductController(service service.ProductService) *productController {
	return &productController{service}
}

func (h *productController) CreateProduct(c *gin.Context) {

	var input entity.ProductInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		//errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("Create Product Failed #CTM001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	input.UserID = c.MustGet("currentUser").(model.User).ID
	var images []entity.ImageInput
	for i, image := range input.Image {
		fmt.Println(i)
		path := fmt.Sprintf("Storage/product/%s", slug.Make(input.Title)+" %"+strconv.Itoa(i)+".png")

		_, err = helper.UploadImage(path, image.ImageName)

		if err != nil {
			responseError := helper.APIResponse("Create Product Failed #CTM019", http.StatusUnsupportedMediaType, "fail", err.Error())
			c.JSON(http.StatusUnsupportedMediaType, responseError)
			return
		}

		image.ImageName = path

		images = append(images, image)

	}

	input.Image = images

	err = h.service.CreateProduct(input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM61", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Product has been created", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)

}

func (h *productController) ListProduct(c *gin.Context) {
	listProduct, err := h.service.ListProduct()
	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM61", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responseError)
		return
	}

	response := helper.APIResponse("Get List product success", http.StatusOK, "success", listProduct)
	c.JSON(http.StatusOK, response)
}

func (h *productController) ListProductUser(c *gin.Context) {
	UserID := c.MustGet("currentUser").(model.User).ID
	listProduct, err := h.service.ListProductUser(UserID)
	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #6", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responseError)
		return
	}

	response := helper.APIResponse("Get List product success", http.StatusOK, "success", listProduct)
	c.JSON(http.StatusOK, response)
}

func (h *productController) UpdateProduct(c *gin.Context) {

	var inputID entity.IdUserInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		responseError := helper.APIResponse("Update Category Failed #RAB0081", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	inputID.User = c.MustGet("currentUser").(model.User)

	var input entity.ProductInput

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("Update Product Failed #CTM001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	var images []entity.ImageInput
	for i, image := range input.Image {
		fmt.Println(i)
		path := fmt.Sprintf("Storage/product/%s", slug.Make(input.Title)+" %"+strconv.Itoa(i)+".png")

		_, err = helper.UploadImage(path, image.ImageName)

		if err != nil {
			responseError := helper.APIResponse("Update Product Failed #CTM019", http.StatusUnsupportedMediaType, "fail", err.Error())
			c.JSON(http.StatusUnsupportedMediaType, responseError)
			return
		}

		image.ImageName = path

		images = append(images, image)
	}

	input.Image = images

	err = h.service.UpdateProduct(inputID, input)

	if err != nil {

		responseError := helper.APIResponse("Update Product Failed #CTM61", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responseError)
		return
	}

	response := helper.APIResponse("Product has been updated", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)

}

func (h *productController) DeleteProduct(c *gin.Context) {
	var inputID entity.IdUserInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		responseError := helper.APIResponse("Delete Product Failed #RAB0091", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	inputID.User = c.MustGet("currentUser").(model.User)

	err = h.service.DeleteProduct(inputID)

	if err != nil {
		responseError := helper.APIResponse("Delete Product Failed #RAB0092", http.StatusBadGateway, "fail", nil)
		c.JSON(http.StatusBadGateway, responseError)
		return
	}
	response := helper.APIResponse("Product has been deleted", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}
