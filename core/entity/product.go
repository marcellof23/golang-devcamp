package entity

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Variants    []Variant `json:"variants"`
	ProductName string    `json:"productName"`
	Description string    `json:"description"`
	Price       *int64    `json:"price,omitempty"`
	Stock       *int64    `json:"stock,omitempty"`
	Discount    *int64    `json:"discount,omitempty"`
}

type ProductInput struct {
	Variants    []VariantInput `json:"variants"`
	ProductName string         `json:"productName"`
	Description string         `json:"description"`
	Price       *int64         `json:"price,omitempty"`
	Stock       *int64         `json:"stock,omitempty"`
	Discount    *int64         `json:"discount,omitempty"`
}
