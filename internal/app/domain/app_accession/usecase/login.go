package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
)

func (u *appAccessionUseCase) Login(ctx context.Context, username string, password string) (*entities.User, error) {
	user, err := u.userSto.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	_, err = u.userSto.ComparePassword(ctx, user.ID, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
