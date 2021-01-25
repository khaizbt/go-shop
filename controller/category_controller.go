package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goshop/entity"
	"goshop/helper"
	"goshop/model"
	"goshop/service"
	"net/http"
)

type categoryController struct {
	service service.CategoryService
}

func CategoryController(service service.CategoryService) *categoryController {
	return &categoryController{service}
}

func (h *categoryController) CreateCategory(c *gin.Context) {
	var input entity.CategoryUserinput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("Create Category Failed #CTG001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	path := fmt.Sprintf("Storage/avatar/%s", input.Slug+".png")

	_, err = helper.UploadImage(path, input.Image)

	if err != nil {
		responsError := helper.APIResponse("Create Product Failed #EMP019", http.StatusUnsupportedMediaType, "fail", err.Error())
		c.JSON(http.StatusUnsupportedMediaType, responsError)
		return
	}

	input.CreatedBy = c.MustGet("currentUser").(model.User).ID

	input.Image = path

	saveCategory, err := h.service.CreateCategory(input)

	if err != nil {
		responseError := helper.APIResponse("Create Category Failed #EMP61", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Category has been created", http.StatusOK, "success", saveCategory)
	c.JSON(http.StatusOK, response)

}

func (h *categoryController) UpdateCategory(c *gin.Context) {
	var inputID entity.IdUserInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		responsError := helper.APIResponse("Update Category Failed #RAT0081", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}
	var input entity.CategoryUserinput

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("Update Category Failed #CTO001", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}
	if input.Image != "" {
		path := fmt.Sprintf("Storage/avatar/%s", input.Slug+".png")

		_, err = helper.UploadImage(path, input.Image)

		if err != nil {
			responsError := helper.APIResponse("Create Product Failed #CTO0192", http.StatusUnsupportedMediaType, "fail", err.Error())
			c.JSON(http.StatusUnsupportedMediaType, responsError)
			return
		}
		input.Image = path
	}
	input.ID = inputID.ID
	input.UpdatedBy = c.MustGet("currentUser").(model.User).ID

	updateCategory, err := h.service.UpdateCategory(input)

	if err != nil {
		responseError := helper.APIResponse("Create Category Failed #CTO0191", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Category has been updated", http.StatusOK, "success", updateCategory)
	c.JSON(http.StatusOK, response)
}
