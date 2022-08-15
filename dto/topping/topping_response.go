package toppingstdo

type ToppingRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Image string `json:"image" form:"id" validate:"required"`
}