package repos

import (
	"errors"
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/data"
	"github.com/s4kibs4mi/movie-pie/log"
	"github.com/s4kibs4mi/movie-pie/services"
	"github.com/s4kibs4mi/movie-pie/utils"
)

type MovieRepo struct {
	dao *data.UserDao
}

var movieRepo *MovieRepo

func NewMovieRepo() *MovieRepo {
	if movieRepo == nil {
		movieRepo = &MovieRepo{}
	}
	return movieRepo
}

func (mr *MovieRepo) Search(s *app.Scope) (interface{}, error) {
	title, ok := s.Ctx.GetQuery("title")

	if !ok {
		return nil, errors.New("title is required")
	}

	pageP, _ := s.Ctx.GetQuery("page")

	page, err := utils.ParseInt(pageP)
	if err != nil {
		log.Log().Errorln(err)
		page = 1
	}

	return services.SearchMovies(title, page)
}
