package models

import (
	"github.com/jinzhu/gorm"
)

type Bill struct {
	gorm.Model
	ProductCode string  `json:"product_code"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	GstAmount   float64 `json:"gst_amount"`
	TotalAmount float64 `json:"total_amount"`
}
