package io

import "mime/multipart"

type ProductDescriptionsWithFileCreateReq struct {
	ProductID uint                  `form:"product_id"`
	Name      string                `form:"name"`
	File      *multipart.FileHeader `form:"descriptions_md"`
}

type ProductDescriptionsWithFileUpdateReq struct {
	Name *string               `form:"name"`
	File *multipart.FileHeader `form:"descriptions_path"`
}
type ProductDescriptionsUpdateReq struct {
	Name string `form:"name"`
	File string `form:"descriptions_path"`
}
type ProductDescriptionsCreateReq struct {
	PublicID         string `json:"public_id"`
	Name             string `json:"name"`
	DescriptionsPath string `json:"path"`
}
