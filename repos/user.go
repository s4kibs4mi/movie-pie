package repos

import (
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/data"
	"github.com/s4kibs4mi/movie-pie/models"
	"time"
)

type UserRepo struct {
	dao *data.UserDao
}

var userRepo *UserRepo

func NewUserRepo() *UserRepo {
	if userRepo == nil {
		userRepo = &UserRepo{}
	}
	return userRepo
}

type reqUserRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (ur *UserRepo) Register(s *app.Scope) (*models.User, error) {
	body := reqUserRegister{}
	if err := s.Ctx.BindJSON(&body); err != nil {
		return nil, err
	}

	u := models.User{}
	u.Name = body.Name
	u.Email = body.Email
	u.Password = body.Password
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	if err := data.NewUserDao().Register(s, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

type reqUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ur *UserRepo) Login(s *app.Scope) (*models.Session, error) {
	body := reqUserLogin{}
	if err := s.Ctx.BindJSON(&body); err != nil {
		return nil, err
	}

	err := data.NewUserDao().CheckLogin(s, body.Email, body.Password)
	if err != nil {
		return nil, err
	}

	ss := models.Session{}
	ss.AccessToken = "1111111111"
	ss.CreateAt = time.Now()
	ss.UpdatedAt = time.Now()

	if err := data.NewUserDao().CreateSession(s, &ss); err != nil {
		return nil, err
	}

	return &ss, nil
}
