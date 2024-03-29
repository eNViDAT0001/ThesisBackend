package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/entities"
	"gorm.io/datatypes"
)

type ListProductInput struct {
	ProductIDs  []uint
	CategoryIDs []uint
	BannerID    uint
	Paging      paging.ParamsInput
}
type ListRecommendedProductInput struct {
	RecommendedProductIDs []uint
	Count                 int
	Paging                paging.ParamsInput
}
type ProductWithQuantities struct {
	entities.Product
	Options datatypes.JSON `json:"options"`
}
type ProductQuantity struct {
	Quantity int    `gorm:"quantity" json:"quantity"`
	Date     string `gorm:"date" json:"date"`
}
