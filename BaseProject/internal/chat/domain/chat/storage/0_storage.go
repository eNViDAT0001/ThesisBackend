package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_query"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/entities"
)

type chatStorage struct {
}

func (c chatStorage) Create(ctx context.Context, input io.MessageInput) (io.MessageInput, error) {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Message{}).Create(&input).Error
	return input, err
}
func (c chatStorage) Update(ctx context.Context, id uint, input io.MessageUpdateInput) error {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Message{}).
		Where("id = ?", id).
		Updates(&input).Error
	return err
}

func (c chatStorage) SeenMessages(ctx context.Context, id uint, userID uint, toID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Model(&entities.Message{}).
		Where("id <= ?", id).
		Where("user_id = ?", userID).
		Where("to_user_id = ?", userID).
		Where("seen = ?", false).
		Update("seen", true).Error
	return err
}

func (c chatStorage) Delete(ctx context.Context, id uint, userID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Message{}.TableName()).
		Where("id = ?", id).
		Where("user_id = ?", userID).
		Delete(entities.Message{}).Error
	return err
}

func (c chatStorage) List(ctx context.Context, input io.ListMessageInput) ([]entities.Message, error) {
	result := make([]entities.Message, 0)
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Message{})

	paging_query.SetPagingQuery(&input.Paging, entities.Message{}.TableName(), query)

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c chatStorage) CountList(ctx context.Context, input io.ListMessageInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Message{})

	paging_query.SetCountListPagingQuery(&input.Paging, entities.Message{}.TableName(), query)

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewChatStorage() chat.Storage {
	return &chatStorage{}
}