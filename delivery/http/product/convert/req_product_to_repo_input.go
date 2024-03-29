package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase/io"
	"github.com/jinzhu/copier"
)

func UpdateProductReqToUpdateProductForm(input *ioHttpHandler.ProductUpdateReq) (ioSto.ProductUpdateForm, error) {
	var result ioSto.ProductUpdateForm
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
func CreateProductReqToCreateProductForm(input *ioHttpHandler.ProductCreateReq) (ioUC.ProductDetailCreateForm, error) {
	var result ioUC.ProductDetailCreateForm

	err := copier.Copy(&result.Product, &input)
	if err != nil {
		return result, err
	}

	err = copier.Copy(&result.Media, &input.Media)
	if err != nil {
		return result, err
	}

	err = copier.Copy(&result.Specifications.Specification, &input.Specification)
	if err != nil {
		return result, err
	}

	err = copier.Copy(&result.Specifications.Options, &input.Options)
	if err != nil {
		return result, err
	}

	err = copier.Copy(&result.Descriptions, &input.Descriptions)
	if err != nil {
		return result, err
	}

	return result, nil
}
