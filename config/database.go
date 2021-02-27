package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goshop/model"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

// SetupDB - setup dabase for sharing to all api
func init() {

	dsn := "root@tcp(127.0.0.1:3306)/go_shop_data?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	err = database.AutoMigrate(&model.User{}, &model.UserType{}, &model.Feature{}, &model.UserTypeFeature{}, &model.Category{}, &model.Product{}, &model.ProductImage{})

	if err != nil {
		panic("failed to migrate")
	}

	var user model.User
	_ = database.First(&user).Error

	if user.ID == 0 {
		err = seed(database)

		if err != nil {
			panic("Failed to run seeds")
		}
	}

	db = database
}
