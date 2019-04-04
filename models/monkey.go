package models

import (
	"github.com/codersgarage/golang-restful-boilerplate/errors"
	"strings"
	"time"
)

type Monkey struct {
	ID        int       `sql:"id;primary key;auto_increment" json:"-"`
	Name      string    `sql:"name" json:"name"`
	Location  string    `sql:"location" json:"location"`
	CreatedAt time.Time `sql:"created_at;not null" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at;not null" json:"updated_at"`
}

func (m *Monkey) TableName() string {
	return "monkeys"
}

func (m *Monkey) Validate() *errors.ValidationError {
	m.Name = strings.TrimSpace(m.Name)
	m.Location = strings.TrimSpace(m.Location)

	ve := errors.ValidationError{}

	if m.Name == "" {
		ve.Add("name", "is required")
	}
	if m.Location == "" {
		ve.Add("location", "is required")
	}

	if len(ve) == 0 {
		return nil
	}
	return &ve
}
