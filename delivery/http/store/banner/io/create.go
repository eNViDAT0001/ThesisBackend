package io

type BannerCreateReq struct {
	UserID     uint   `json:"user_id"`
	Title      string `json:"title"`
	Collection string `json:"collection"`
	Discount   int    `json:"discount"`
	Image      string `json:"image"`
	EndTime    string `json:"end_time"`
	ProductIDs []uint `json:"product_ids"`
}
