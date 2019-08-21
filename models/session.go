package models

import "time"

type Session struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreateAt     time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (s *Session) TableName() string {
	return "sessions"
}
