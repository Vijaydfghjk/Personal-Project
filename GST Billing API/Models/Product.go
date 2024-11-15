package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ProductCode  string  `json:"product_code" gorm:"unique;not null"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductGst   float64 `json:"product_gst"`
}
