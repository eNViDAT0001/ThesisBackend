package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"
	"gorm.io/gorm"
)

func (u providerUseCase) ListProvider(ctx context.Context, filter paging.ParamsInput) (providers []entities.Provider, total int64, err error) {
	total, err = u.providerSto.CountListProvider(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}
	providers, err = u.providerSto.ListProvider(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return providers, total, err
}
