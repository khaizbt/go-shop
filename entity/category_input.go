package entity

type (
	CategoryUserinput struct {
		ID        int
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Image     string `json:"image"`
		CreatedBy int    `json:"-"`
		UpdatedBy int    `json:"-"`
		DeletedBy int    `json:"-"`
	}

	IdUserInput struct {
		ID int `uri:"id" binding:"required"` //Ambil id dari URL
	}
)