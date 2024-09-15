package adapter

import (
	"main/src/chat/domain/model"
	"main/src/chat/domain/repository"

	fiberlog "github.com/gofiber/fiber/v2/log"

	"github.com/jmoiron/sqlx"
)

type ChatPSQLRepository struct {
	client *sqlx.DB
}

func NewChatPSQLRepository(client *sqlx.DB) repository.ChatRepository {
	return &ChatPSQLRepository{client: client}
}

func (r *ChatPSQLRepository) CreateChat(chat *model.Chat) error {
	// Use the RETURNING clause to get the generated ID
	query := `
		INSERT INTO chats (user_id, title, created_at, system_fingerprint, model_used, total_tokens)
		VALUES (:user_id, :title, :created_at, :system_fingerprint, :model_used, :total_tokens)
		RETURNING id
	`

	// Execute the query and get the new ID
	rows, err := r.client.NamedQuery(query, chat)
	if err != nil {
		fiberlog.Infof("Error creating chat: %v", err)
		return err
	}
	defer rows.Close()

	// Retrieve the ID from the result
	if rows.Next() {
		if err := rows.Scan(&chat.ID); err != nil {
			fiberlog.Infof("Error scanning returned ID: %v", err)
			return err
		}
	}

	return nil
}



func (r *ChatPSQLRepository) GetChatByID(id int) (*model.Chat, error) {
	var chat model.Chat
	err := r.client.Get(&chat, "SELECT * FROM chats WHERE id = $1", id)
	if err != nil {
		fiberlog.Infof("Error getting chat by ID: %v", err)
		return nil, err
	}
	return &chat, nil
}

func (r *ChatPSQLRepository) GetAllChats() ([]model.Chat, error) {
	var chats []model.Chat
	err := r.client.Select(&chats, "SELECT * FROM chats")
	if err != nil {
		fiberlog.Infof("Error getting all chats: %v", err)
		return nil, err
	}
	return chats, nil
}

func (r *ChatPSQLRepository) UpdateChat(chat *model.Chat) error {
	_, err := r.client.NamedExec(`
		UPDATE chats
		SET user_id = :user_id, title = :title, created_at = :created_at, system_fingerprint = :system_fingerprint, model_used = :model_used, total_tokens = :total_tokens
		WHERE id = :id
	`, chat)
	if err != nil {
		fiberlog.Infof("Error updating chat: %v", err)
		return err
	}
	return nil
}

func (r *ChatPSQLRepository) DeleteChat(id int) error {
	_, err := r.client.Exec("DELETE FROM chats WHERE id = $1", id)
	if err != nil {
		fiberlog.Infof("Error deleting chat: %v", err)
		return err
	}
	return nil
}
