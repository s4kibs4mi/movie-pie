package data

import (
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/models"
)

type UserDao struct {
}

var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{}
	}
	return userDao
}

func (ud *UserDao) Register(s *app.Scope, u *models.User) error {
	return s.DB.Create(u).Error
}

func (ud *UserDao) CreateSession(s *app.Scope, ss *models.Session) error {
	return s.DB.Create(ss).Error
}

func (ud *UserDao) CheckLogin(s *app.Scope, email, password string) error {
	ss := models.Session{}
	if err := s.DB.Table(ss.TableName()).Where("email = ? AND password = ?", email, password).Find(&ss).Error; err != nil {
		return err
	}
	return nil
}
