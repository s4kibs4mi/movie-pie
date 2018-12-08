package models

import (
	"github.com/codersgarage/golang-restful-boilerplate/errors"
	"strings"
	"time"
)

type Monkey struct {
	ID        int       `sql:"id;primary key;auto_increment" json:"-"`
	UUID      string    `sql:"uuid;not null;unique" json:"uuid"`
	Name      string    `sql:"name" json:"name"`
	CreatedAt time.Time `sql:"created_at;not null" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at;not null" json:"updated_at"`
}

func (h *Monkey) TableName() string {
	return "monkeys"
}

func (h *Monkey) Validate() *errors.ValidationError {
	h.Name = strings.TrimSpace(h.Name)

	ve := errors.ValidationError{}

	if h.Name == "" {
		ve.Add("name", "is required")
	}

	if len(ve) == 0 {
		return nil
	}

	return &ve
}
