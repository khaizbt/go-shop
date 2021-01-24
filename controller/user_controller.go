package controller

import (
	"fmt"
	"goshop/config"
	"goshop/entity"
	"goshop/helper"
	"goshop/model"
	"goshop/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	userController struct {
		userService service.UserService
		authService config.AuthService
	}
)

func NewUserController(userService service.UserService, authService config.AuthService) *userController {
	return &userController{userService, authService}
}

type UserFormatter struct {
	UserID int     `json:"id"`
	Email  string  `json:"email"`
	Phone  *string `json:"phone"`
	Token  string  `json:"token"`
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

	path := fmt.Sprintf("Storage/avatar/%s", input.Username+".png")

	_, err = helper.UploadImage(path, input.Avatar)

	if err != nil {
		responsError := helper.APIResponse("Create Product Failed #EMP019", http.StatusUnsupportedMediaType, "fail", err.Error())
		c.JSON(http.StatusUnsupportedMediaType, responsError)
		return
	}

	input.Avatar = path

	createUser, err := h.userService.CreateUser(input)

	if err != nil {
		responsError := helper.APIResponse("Create Account Failed", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responsError)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", createUser)
	c.JSON(http.StatusOK, response)
}

func (h *userController) ListUser(c *gin.Context) {
	var input entity.DataUserInput

	name := c.Query("name")
	email := c.Query("email")
	phone := c.Query("phone")
	username := c.Query("username")
	userType, _ := strconv.Atoi(c.Query("user_type"))

	input.Name = name
	input.Email = email
	input.Phone = phone
	input.Username = username
	input.UserTypeID = userType

	listUser, err := h.userService.ListUser(input)
	if err != nil {
		responseError := helper.APIResponse("Get List User Failed #JAH002", http.StatusBadGateway, "fail", err.Error())
		c.JSON(http.StatusBadGateway, responseError)
		return
	}

	response := helper.APIResponse("Get List User Success", http.StatusOK, "success", listUser)
	c.JSON(http.StatusOK, response)
}
