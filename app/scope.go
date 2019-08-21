package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Scope struct {
	Ctx *gin.Context
	DB  *gorm.DB
}

func NewScope(db *gorm.DB, ctx *gin.Context) *Scope {
	return &Scope{
		DB:  db,
		Ctx: ctx,
	}
}
