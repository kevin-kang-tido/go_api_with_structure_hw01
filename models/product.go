package models

import (
    "gorm.io/gorm"
)

type Product struct {

	gorm.Model // handle createAt  updateAt and deleteAt
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
    Price       float64 `json:"price" binding:"required"`
	Qty uint `json:"qty"`

}