package application

import "main/src/chat/domain/model"

type ChatService interface {
	CreateChat(chat *model.Chat) error
	GetChatByID(id int) (*model.Chat, error)
	GetAllChats() ([]model.Chat, error)
	UpdateChat(chat *model.Chat) error
	DeleteChat(id int) error
}
