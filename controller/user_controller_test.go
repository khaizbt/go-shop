package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"golang.org/x/crypto/bcrypt"
	"goshop/config"
	"goshop/entity"
	"goshop/middleware"
	"goshop/repository"
	"goshop/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var userRepo = repository.NewUserRepository()
var userService = service.NewUserService(userRepo)
var authService = config.NewServiceAuth()
var userTestController = NewUserController(userService, authService)

func TestUserController_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	input := entity.LoginEmailInput{
		Email:    "khaiz@ggggg.com",
		Password: "larashop",
	}
	requestBody, _ := json.Marshal(input)
	r := gin.Default()
	r.POST("/api/v1/login", userTestController.Login)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	//r.Run(":8000", w)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestUserController_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	input := entity.DataUserInput{
		Name:       "Tammam Testing",
		Email:      "tamam@mailinator.com",
		Username:   "tamam",
		UserTypeID: 1,
	}

	//Create Password
	password, _ := bcrypt.GenerateFromPassword([]byte("654321"), bcrypt.MinCost)
	//
	////Check Password
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte("654321"))

	assert.Equal(t, nil, err)

	input.Password = string(password)

	requestBody, _ := json.Marshal(input)

	r := gin.Default()
	r.POST("/api/v1/create", middleware.AuthMiddlewareUser(authService, userService, 1), userTestController.CreateUser)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/create", bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6MiwidXNlcl9pZCI6MX0.kOR5FGpvbacvjjUVEeyKS8nOCDNHynb8Fn-toZNSaVo")
	w := httptest.NewRecorder()
	//r.Run(":8000", w)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
