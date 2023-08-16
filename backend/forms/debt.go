package forms

import "time"

type Dept struct {
	ID 							uint64 			`gorm:"primary_key;auto_increment" json:"id"`
	Title 					string 			`json:"title" binding:"required"`
	Description 		string 			`json:"description"`
	Price 					float64 		`json:"price" binding:"required"`
	CreatedAt 			*time.Time 	`json:"created_at" binding:"required"`
}
