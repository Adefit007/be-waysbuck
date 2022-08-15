package toppingstdo

type UpdateTopping struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" gorm:"type: int" form:"price"`
	Image string `json:"image" form:"image"`
}