package entity

import (
	"github.com/jinzhu/gorm"
)

type Variant struct {
	gorm.Model
	VariantName string `json:"variantName"`
	ProductID   uint
	Price       int64 `json:"price"`
	Stock       int64 `json:"stock"`
}

type VariantInput struct {
	VariantName string `json:"variantName"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}
