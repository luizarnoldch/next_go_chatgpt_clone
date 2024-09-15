package model

import "time"

type Message struct {
	ID               int       `json:"id" db:"id"`
	ChatID           int       `json:"chat_id" db:"chat_id"`
	Role             string    `json:"role" db:"role"`
	Content          string    `json:"content" db:"content"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	FinishReason     string    `json:"finish_reason,omitempty" db:"finish_reason"`
	PromptTokens     int       `json:"prompt_tokens" db:"prompt_tokens"`
	CompletionTokens int       `json:"completion_tokens" db:"completion_tokens"`
	TotalTokens      int       `json:"total_tokens" db:"total_tokens"`
}
