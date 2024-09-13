package application

import (
	"main/src/chat/domain/model"
	"main/src/chat/domain/repository"
	"time"
)

type ChatPSQLServie struct {
	r repository.ChatRepository
}

func NewChatPSQLServie(repositoty repository.ChatRepository) ChatService {
	return &ChatPSQLServie{
		r: repositoty,
	}
}

func (s ChatPSQLServie) CreateChat(chat *model.Chat) error {

	if chat.CreatedAt.IsZero() {
		chat.CreatedAt = time.Now()
	}

	return s.r.CreateChat(chat)
}

func (s ChatPSQLServie) GetChatByID(id int) (*model.Chat, error) {
	return s.r.GetChatByID(id)
}

func (s ChatPSQLServie) GetAllChats() ([]model.Chat, error) {
	return s.r.GetAllChats()
}

func (s ChatPSQLServie) UpdateChat(chat *model.Chat) error {
	return s.r.UpdateChat(chat)
}

func (s ChatPSQLServie) DeleteChat(id int) error {
	return s.r.DeleteChat(id)
}
