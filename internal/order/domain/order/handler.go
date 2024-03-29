package order

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	ListByUserID() func(ctx *gin.Context)
	ListPreviewByUserID() func(ctx *gin.Context)
	ListByProviderID() func(ctx *gin.Context)
	ListPreviewByProviderID() func(ctx *gin.Context)

	List() func(ctx *gin.Context)
	ListPreview() func(ctx *gin.Context)

	ListReport() func(ctx *gin.Context)

	GetByOrderID() func(ctx *gin.Context)

	CreateOrder() func(ctx *gin.Context)

	VerifyDeliveredStatus() func(ctx *gin.Context)
	UpdateOrderStatus() func(ctx *gin.Context)
	UpdateOrder() func(ctx *gin.Context)
	UpdateOrderPayment() func(ctx *gin.Context)
	CancelOrder() func(ctx *gin.Context)
	DeleteOrder() func(ctx *gin.Context)

	RemoveInvalidOrder() error
}
