package config

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"goshop/model"
)

func seed(db *gorm.DB) error {
	userTypes := []model.UserType{
		{ID: 1, Name: "Super Admin"},
		{ID: 2, Name: "Admin"},
		{ID: 3, Name: "User"},
	}

	for _, userType := range userTypes {

		err := db.Create(&userType).Error
		if err != nil {
			return err
		}
	}

	users := []model.User{
		{ID: 1, Name: "Super Admin", Email: "superadmin@example.com", Username: "superadmin", IDUserType: 1, Address: "Yogyakarta", Status: "Active"},
		{ID: 2, Name: "Admin", Email: "admin@example.com", Username: "admin", IDUserType: 2, Address: "Yogyakarta", Status: "Active"},
		{ID: 3, Name: "user", Email: "user@example.com", Username: "user", IDUserType: 3, Address: "Yogyakarta", Status: "Active"},
	}

	for _, user := range users {
		password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.MinCost)
		user.Password = string(password)
		err := db.Create(&user).Error
		if err != nil {
			return err
		}
	}

	return nil
}
