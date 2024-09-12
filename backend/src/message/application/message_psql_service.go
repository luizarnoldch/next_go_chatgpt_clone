package application

import (
	"main/src/message/domain/model"
	"main/src/message/domain/repository"
)

type MessagePSQLServie struct {
	r repository.MessageRepository
}

func NewMessagePSQLServie(repositoty repository.MessageRepository) MessageService {
	return &MessagePSQLServie{
		r: repositoty,
	}
}

func (s MessagePSQLServie) CreateMessage(message *model.Message) error {
	return s.r.CreateMessage(message)
}
func (s MessagePSQLServie) GetMessageByID(id int) (*model.Message, error) {
	return s.r.GetMessageByID(id)
}
func (s MessagePSQLServie) GetAllMessages() ([]model.Message, error) {
	return s.r.GetAllMessages()
}
func (s MessagePSQLServie) UpdateMessage(message *model.Message) error {
	return s.r.UpdateMessage(message)
}
func (s MessagePSQLServie) DeleteMessage(id int) error {
	return s.r.DeleteMessage(id)
}
