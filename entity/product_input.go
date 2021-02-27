package entity

type (
	ProductInput struct {
		Title            string `json:"title" binding:"required"`
		Description      string `json:"description" binding:"required"`
		Address          string `json:"address" binding:"required"`
		Phone            string `json:"phone" binding:"required"`
		InstagramAccount string `json:"instagram_account" binding:"required"`
		PurchasePrice    int    `json:"purchase_price" binding:"required"`
		SellingPrice     int    `json:"selling_price" binding:"required"`
		CategoryID       int    `json:"category_id" binding:"required"`
		UserID           int
		Image            []ImageInput `json:"image"`
	}

	ImageInput struct {
		ImageName string `json:"image_name"`
		IsPrimary bool   `json:"is_primary"`
	}
)
