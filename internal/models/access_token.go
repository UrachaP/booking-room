package models

import "time"

type AccessTokens struct {
	ID        int       `json:"id"`
	IDToken   string    `json:"id_token"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Revoked   string    `json:"revoked" gorm:"default:"0"" `
}

type LoginHistory struct {
	FullName  string    `json:"full_name"`
	IDToken   string    `json:"id_token"`
	CreatedAt time.Time `json:"created_at"`
	Revoked   string    `json:"revoked"`
}

func (LoginHistory) TableName() string {
	return "access_tokens"
}
