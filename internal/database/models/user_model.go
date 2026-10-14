package models

type UserModel struct {
	ID         int    `db:"id" json:"id"`
	Username   string `db:"username" json:"username" validate:"required"`
	Email      string `db:"email" json:"email" validate:"required"`
	Password   string `db:"password" json:"password" validate:"required"`
	IsVerified bool   `db:"is_verified" json:"is_verified"`
	CreatedAt  string `db:"created_at" json:"created_at"`
	UpdatedAt  string `db:"updated_at" json:"updated_at"`
}
