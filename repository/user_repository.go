package repository

import (
	"goshop/config"
	"goshop/entity"
	"goshop/model"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindUserByEmail(email string) (model.User, error)
		FindByID(ID int) (model.User, error)
		UpdateProfile(user model.User) (model.User, error)
		CreateUser(user model.User) (model.User, error)
		UserFeature(feature model.UserTypeFeature) (model.UserTypeFeature, error)
		ListUser(input entity.DataUserInput) ([]model.User, error)
	}

	repository struct {
		db *gorm.DB
	}
)

func NewUserRepository() *repository {
	return &repository{config.GetDB()}
}

func (r *repository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (model.User, error) {
	var user model.User

	err := r.db.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateProfile(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UserFeature(feature model.UserTypeFeature) (model.UserTypeFeature, error) {
	err := r.db.Where("id_user_type = ?", feature.IDUserType).Where("id_feature = ?", feature.IDFeature).Find(&feature).Error

	if err != nil {
		return feature, err
	}

	return feature, nil
}

func (r *repository) ListUser(input entity.DataUserInput) ([]model.User, error) {
	var users []model.User
	user := r.db.Preload("UserType")

	if input.Name != "" {
		user.Where("name LIKE ?", "%"+input.Name+"%")
	}

	if input.Username != "" {
		user.Where("username LIKE ?", "%"+input.Username+"%")
	}

	if input.Phone != "" {
		user.Where("phone = ?", input.Phone)
	}

	if input.Email != "" {
		user.Where("email LIKE ?", "%"+input.Email+"%")
	}

	if input.UserTypeID != 0 {
		user.Where("id_user_type = ?", input.UserTypeID)
	}

	err := user.Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
