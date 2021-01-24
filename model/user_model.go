package model

import (
	"time"
)

type (
	User struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		Password   string    `json:"-"`
		Username   string    `json:"username"`
		IDUserType int       `json:"user_type_id"`
		Address    string    `json:"address"`
		Phone      *string   `json:"phone"`
		Avatar     string    `json:"avatar"`
		Status     string    `json:"-"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		UserType   UserType  `json:"user_type" gorm:"foreignKey:IDUserType"`
	}

	UserType struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Feature struct {
		ID        int    `json:"id"`
		Key       string `json:"key"`
		Value     string `json:"value"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	UserTypeFeature struct {
		ID         int `json:"id"`
		IDUserType int `json:"user_type_id"`
		IDFeature  int `json:"feature_id"`
	}
)
