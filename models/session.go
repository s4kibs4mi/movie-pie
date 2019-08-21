package models

import "time"

type Session struct {
	ID           int64     `json:"id" sql:"primary_key;auto_increment"`
	AccessToken  string    `json:"access_token" sql:"index;unique"`
	RefreshToken string    `json:"refresh_token" sql:"index;unique"`
	CreateAt     time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (s *Session) TableName() string {
	return "sessions"
}
