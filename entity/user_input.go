package entity

type (
	LoginEmailInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	DataUserInput struct {
		ID         int
		Name       string `json:"name" binding:"required"`
		Email      string `json:"email" binding:"required"`
		Username   string `json:"username" binding:"required"`
		Address    string `json:"address"`
		Phone      int    `json:"phone"`
		Password   string `json:"password" binding:"required"`
		Avatar     string `json:"avatar"`
		UserTypeID int    `json:"user_type_id" binding:"required"`
	}
)
