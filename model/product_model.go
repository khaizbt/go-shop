package model

import "time"

type (
	Product struct {
		ID               int            `json:"id"`
		CategoryId       int            `json:"category_id"`
		Title            string         `json:"title"`
		Slug             string         `json:"slug"`
		Description      string         `json:"description"`
		Address          string         `json:"address"`
		Phone            string         `json:"phone"`
		InstagramAccount string         `json:"instagram_account"`
		PurchasePrice    int            `json:"purchase_price"`
		SellingPrice     int            `json:"selling_price"`
		UserID           int            `json:"user_id"`
		CreatedAt        time.Time      `json:"created_at"`
		UpdatedAt        time.Time      `json:"updated_at"`
		Category         Category       `gorm:"foreignKey:CategoryId" json:"category,omitempty"`
		User             User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
		ProductImage     []ProductImage `json:"product_image,omitempty"`
	}

	ProductImage struct {
		ID        int       `json:"id"`
		ProductID int       `json:"product_id"`
		ImageName string    `json:"image_name"`
		IsPrimary bool      `json:"is_primary"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	}
)
