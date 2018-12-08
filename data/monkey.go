package data

import (
	"github.com/codersgarage/golang-restful-boilerplate/app"
	"github.com/codersgarage/golang-restful-boilerplate/models"
)

type MonkeyDao struct {
}

func NewMonkeyDao() *MonkeyDao {
	return &MonkeyDao{}
}

func (hd *MonkeyDao) SaveMonkey(scope *app.Scope, monkey *models.Monkey) error {
	return scope.DB.Table(monkey.TableName()).Create(monkey).Error
}

func (hd *MonkeyDao) GetMonkey(scope *app.Scope, UUID string, v *models.Monkey) error {
	return scope.DB.
		Table(v.TableName()).
		Find(v, "uuid = ?", UUID).Error
}

func (hd *MonkeyDao) ListMonkeys(scope *app.Scope, v interface{}) error {
	monkey := models.Monkey{}
	return scope.DB.
		Table(monkey.TableName()).
		Find(v).Error
}
