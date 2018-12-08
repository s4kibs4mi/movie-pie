package app

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Scope struct {
	Request *http.Request
	DB      *gorm.DB
}

func NewScope(db *gorm.DB, r *http.Request) *Scope {
	return &Scope{
		DB:      db,
		Request: r,
	}
}
