package application

import "main/src/message/domain/model"

type MessageService interface {
	CreateMessage(message *model.Message) error
	GetMessageByID(id int) (*model.Message, error)
	GetAllMessages() ([]model.Message, error)
	UpdateMessage(message *model.Message) error
	DeleteMessage(id int) error
}
