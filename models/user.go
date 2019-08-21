package models

import "time"

type User struct {
	ID        int64     `json:"id" sql:"primary_key;auto_increment"`
	Name      string    `json:"name"`
	Email     string    `json:"email" sql:"index;unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
