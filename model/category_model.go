package model

import "time"

type (
	Category struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Slug      *string   `json:"slug"`
		Image     *string   `json:"image"`
		CreatedBy int       `json:"-"`
		UpdatedBy *int      `json:"-"`
		DeletedBy *int      `json:"-"`
		DeletedAt time.Time `json:"-"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
