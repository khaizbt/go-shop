package controller

import (
	"goshop/config"
	"goshop/entity"
	"goshop/helper"
	"goshop/model"
	"goshop/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
	authService config.AuthService
}

func NewUserController(userService service.UserService, authService config.AuthService) *userController {
	return &userController{userService, authService}
}

type UserFormatter struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	Phone  int    `json:"phone"`
	Token  string `json:"token"`
}

func FormatUser(user model.User, token string) UserFormatter { //Token akan didapatkan dari JWT
	formatter := UserFormatter{
		UserID: user.ID,
		Email:  user.Email,
		Phone:  user.Phone,
		Token:  token,
	}

	return formatter
}

func (h *userController) Login(c *gin.Context) {
	var input entity.LoginEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		responseError := helper.APIResponse("Login Failed #LOG001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		responseError := helper.APIResponse("Login Failed #LOG002", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}
	token, err := h.authService.GenerateTokenUser(loggedInUser.ID, loggedInUser.IDUserType)
	if err != nil {
		responsError := helper.APIResponse("Login Failed", http.StatusBadGateway, "fail", "Unable to generate token")
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	response := helper.APIResponse("Login Success", http.StatusOK, "success", FormatUser(loggedInUser, token))

	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateProfile(c *gin.Context) {
	var input entity.DataUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}

		responseError := helper.APIResponse("Create Account Failed", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	input.ID = c.MustGet("currentUser").(model.User).ID
	updateUser, err := h.userService.UpdateProfile(input)
	if err != nil {
		responsError := helper.APIResponse("Create Account Failed", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", updateUser)
	c.JSON(http.StatusOK, response)
}

func (h *userController) CreateUser(c *gin.Context) {
	var input entity.DataUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidationError(err)}
		responseError := helper.APIResponse("Create Account Failed", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	createUser, err := h.userService.CreateUser(input)

	if err != nil {
		responsError := helper.APIResponse("Create Account Failed", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", createUser)
	c.JSON(http.StatusOK, response)
}
