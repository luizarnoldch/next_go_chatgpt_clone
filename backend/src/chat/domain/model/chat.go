package model

type Chat struct {
	ID                int    `json:"id" db:"id"`
	UserID            int    `json:"user_id,omitempty" db:"user_id"`
	CreatedAt         string `json:"created_at" db:"created_at"`
	SystemFingerprint string `json:"system_fingerprint,omitempty" db:"system_fingerprint"`
	ModelUsed         string `json:"model_used,omitempty" db:"model_used"`
	TotalTokens       int    `json:"total_tokens" db:"total_tokens"`
}
