package productsdto

type CreateProduct struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" gorm:"type: int" form:"price" validate:"required"`
	Image string `json:"image" form:"image"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" gorm:"type: int" form:"price"`
	Image string `json:"image" form:"image"`
}
