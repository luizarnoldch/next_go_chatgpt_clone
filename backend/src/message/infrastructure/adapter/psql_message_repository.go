package repository

import (
	"main/src/message/domain/model"
	"main/src/message/domain/repository"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
)

type MessagePSQLRepository struct {
	client *sqlx.DB
}

func NewMessagePSQLRepository(client *sqlx.DB) repository.MessageRepository {
	return &MessagePSQLRepository{client: client}
}

func (r *MessagePSQLRepository) CreateMessage(message *model.Message) error {
	_, err := r.client.NamedExec(`
		INSERT INTO messages (chat_id, role, content, created_at, finish_reason, prompt_tokens, completion_tokens, total_tokens)
		VALUES (:chat_id, :role, :content, :created_at, :finish_reason, :prompt_tokens, :completion_tokens, :total_tokens)
	`, message)
	if err != nil {
		fiberlog.Infof("Error creating message: %v", err)
		return err
	}
	return nil
}

func (r *MessagePSQLRepository) GetMessageByID(id int) (*model.Message, error) {
	var message model.Message
	err := r.client.Get(&message, "SELECT * FROM messages WHERE id = $1", id)
	if err != nil {
		fiberlog.Infof("Error getting message by ID: %v", err)
		return nil, err
	}
	return &message, nil
}

func (r *MessagePSQLRepository) GetAllMessages() ([]model.Message, error) {
	var messages []model.Message
	err := r.client.Select(&messages, "SELECT * FROM messages")
	if err != nil {
		fiberlog.Infof("Error getting all messages: %v", err)
		return nil, err
	}
	return messages, nil
}

func (r *MessagePSQLRepository) UpdateMessage(message *model.Message) error {
	_, err := r.client.NamedExec(`
		UPDATE messages
		SET chat_id = :chat_id, role = :role, content = :content, created_at = :created_at, finish_reason = :finish_reason,
			prompt_tokens = :prompt_tokens, completion_tokens = :completion_tokens, total_tokens = :total_tokens
		WHERE id = :id
	`, message)
	if err != nil {
		fiberlog.Infof("Error updating message: %v", err)
		return err
	}
	return nil
}

func (r *MessagePSQLRepository) DeleteMessage(id int) error {
	_, err := r.client.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		fiberlog.Infof("Error deleting message: %v", err)
		return err
	}
	return nil
}
