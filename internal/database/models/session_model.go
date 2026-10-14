package models

type SessionModel struct {
	ID               int    `db:"id" json:"id"`
	UserID           int    `db:"user_id" json:"user_id" validate:"required"`
	RefreshTokenHash string `db:"refresh_token_hash" json:"refresh_token_hash" validate:"required"`
	IsRevoked        bool   `db:"is_revoked" json:"is_revoked"`
	CreatedAt        string `db:"created_at" json:"created_at"`
	UpdatedAt        string `db:"updated_at" json:"updated_at"`
}
