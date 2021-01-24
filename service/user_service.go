package service

import (
	"errors"
	"goshop/entity"
	"goshop/model"
	"goshop/repository"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Login(input entity.LoginEmailInput) (model.User, error)
		GetUserById(ID int) (model.User, error)
		UpdateProfile(input entity.DataUserInput) (bool, error)
		CheckFeature(idUserType int, idFeature int) (model.UserTypeFeature, error)
		CreateUser(input entity.DataUserInput) (bool, error)
	}

	service struct {
		repository repository.UserRepository
	}
)

func NewUserService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) GetUserById(ID int) (model.User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID != ID {
		return user, errors.New("User Tidak Ditemukan")
	}

	return user, nil
}

func (s *service) Login(input entity.LoginEmailInput) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	// if user.StatusRegistered == 0 { //Handle if not verified
	// 	return user, errors.New("Verified Your Account and Try again")
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateProfile(input entity.DataUserInput) (bool, error) {
	user, err := s.repository.FindByID(input.ID)

	if err != nil {
		return false, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Username = input.Username

	_, err = s.repository.UpdateProfile(user)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) CheckFeature(idUserType int, idFeature int) (model.UserTypeFeature, error) {
	feature := model.UserTypeFeature{
		IDUserType: idUserType,
		IDFeature:  idFeature,
	}

	feature, err := s.repository.UserFeature(feature)

	if err != nil {
		return feature, err
	}

	return feature, nil
}

func (s *service) CreateUser(input entity.DataUserInput) (bool, error) {
	cekUser, err := s.repository.FindUserByEmail(input.Email)

	if cekUser.ID != 0 {
		return false, errors.New("Email has been registered")
	}

	user := model.User{
		Email:      input.Email,
		Username:   input.Username,
		Name:       input.Name,
		IDUserType: input.UserTypeID,
		Address:    input.Address,
		Phone:      &input.Phone,
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return false, err
	}

	user.Password = string(password)

	_, err = s.repository.CreateUser(user)

	if err != nil {
		return false, err
	}

	return true, nil
}
