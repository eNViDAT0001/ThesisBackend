package chat

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type Storage interface {
	Create(ctx context.Context, input io.MessageInput) (io.MessageInput, error)
	Update(ctx context.Context, id uint, input io.MessageUpdateInput) error
	SeenMessages(ctx context.Context, id uint, userID uint, toID uint) error
	Delete(ctx context.Context, id uint, userID uint) error
	List(ctx context.Context, input io.ListMessageInput) ([]entities.Message, error)
	CountList(ctx context.Context, input io.ListMessageInput) (int64, error)
}