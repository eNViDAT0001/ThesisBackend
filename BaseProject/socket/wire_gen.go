// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package socket

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/chat"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/notify"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage"
	"github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/usecase"
	storage2 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage"
	usecase2 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/usecase"
)

// Injectors from socket_wire.go:

func initSocketCollection() *WebSocketCollection {
	chatStorage := storage.NewChatStorage()
	useCase := usecase.NewChatUseCase(chatStorage)
	chatWebSocketHandler := chat.NewSocketChatHandler(useCase)
	notificationStorage := storage2.NewNotificationStorage()
	notificationUseCase := usecase2.NewNotificationUseCase(notificationStorage)
	notifyWebSocketHandler := notify.NewSocketNotificationHandler(notificationUseCase)
	webSocketCollection := NewSocketCollection(chatWebSocketHandler, notifyWebSocketHandler)
	return webSocketCollection
}