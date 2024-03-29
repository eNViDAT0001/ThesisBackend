package order

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage/io"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/order/entities"
)

type Storage interface {
	ListByUserID(ctx context.Context, userID uint, input paging.ParamsInput) ([]entities.Order, error)
	CountListByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (total int64, err error)
	ListPreviewByUserID(ctx context.Context, userID uint, input paging.ParamsInput) ([]io.OrderPreview, error)
	CountPreviewByUserID(ctx context.Context, userID uint, input paging.ParamsInput) (total int64, err error)

	ListUnPayOrder(ctx context.Context) ([]entities.Order, error)
	ListUnConfirmedDeliveredOrder(ctx context.Context) ([]entities.Order, error)

	ListByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) ([]entities.Order, error)
	CountByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (total int64, err error)

	ListPreviewByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) ([]io.OrderPreview, error)
	CountPreviewByProviderID(ctx context.Context, providerID uint, input paging.ParamsInput) (total int64, err error)

	ListQuantity(ctx context.Context, input paging.ParamsInput) ([]io.OrderReportQuantity, error)
	ListQuantityCount(ctx context.Context, input paging.ParamsInput) (int64, error)

	CountList(ctx context.Context, input paging.ParamsInput) (total int64, err error)
	List(ctx context.Context, input paging.ParamsInput) ([]entities.Order, error)
	CountPreview(ctx context.Context, input paging.ParamsInput) (total int64, err error)
	ListPreview(ctx context.Context, input paging.ParamsInput) ([]io.OrderPreview, error)

	GetByOrderID(ctx context.Context, orderID uint) (entities.Order, error)

	CreateOrder(ctx context.Context, order io.CreateOrderForm, items []io2.CreateOrderItemForm, cartItemsIDs []uint, couponCodes []string) (createdOrders []io.CreateOrderForm, err error)

	UpdateOrderStatus(ctx context.Context, orderID uint, status entities.OrderStatus) error
	UpdateOrder(ctx context.Context, orderID uint, input io.UpdateOrderForm) error
	UpdateOrderPayment(ctx context.Context, orderIDs []uint, input io.UpdateOrderPaymentForm) error
	CancelOrder(ctx context.Context, orderID uint) error
	DeleteOrder(ctx context.Context, orderID uint) error
	DeleteOrders(ctx context.Context, ids []uint) error

	GetOrdersReport(ctx context.Context, filter io.OrderReportFilter) (report io.OrderReport, err error)
	UpdateDeliveredOrderStatus(ctx context.Context, id uint, image string) error
	VerifyDeliveredOrder(ctx context.Context, orderID uint, userID uint) error
}
