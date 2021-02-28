package entity

import "goshop/model"

type (
	CategoryUserinput struct {
		ID        int
		Name      string `json:"name"`
		Image     string `json:"image"`
		CreatedBy int    `json:"-"`
		UpdatedBy int    `json:"-"`
		DeletedBy int    `json:"-"`
	}

	IdUserInput struct {
		ID   int `uri:"id" binding:"required"` //Ambil id dari URL
		User model.User
	}
)
