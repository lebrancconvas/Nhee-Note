package forms

type Dept struct {
	ID uint64 `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price float64 `json:"price" binding:"required"`
	CreatedAt string `json:"created_at" binding:"required"`
}
