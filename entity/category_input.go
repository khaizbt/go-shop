package entity

type (
	CategoryUserinput struct {
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Image     string `json:"image"`
		CreatedBy int    `json:"-"`
		UpdatedBy int    `json:"-"`
		DeletedBy int    `json:"-"`
	}
)
