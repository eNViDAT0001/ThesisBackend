package provider

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
)

type Storage interface {
	CreateProvider(ctx context.Context, input io.ProviderForm) (io.ProviderForm, error)
	GetProviderByID(ctx context.Context, provider uint) (entities.Provider, error)
	UpdateProvider(ctx context.Context, providerID uint, input io.ProviderUpdateForm) error
	DeleteProviderByIDs(ctx context.Context, providerID []uint) error
	ListProviderByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) ([]entities.Provider, error)
	ListProvider(ctx context.Context, filter paging.ParamsInput) ([]entities.Provider, error)
	ListProviderQuantity(ctx context.Context, filter paging.ParamsInput) ([]io.ProviderQuantity, error)
	CountListProvider(ctx context.Context, filter paging.ParamsInput) (total int64, err error)
	CountListProviderByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (total int64, err error)
	CountListProviderQuantity(ctx context.Context, filter paging.ParamsInput) (total int64, err error)
	GetProviderFullDetailByID(ctx context.Context, id uint) (io.ProviderFullDetail, error)
}
