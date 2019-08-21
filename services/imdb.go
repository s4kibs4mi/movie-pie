package services

import (
	"fmt"
	"github.com/nahid/gohttp"
	"github.com/s4kibs4mi/movie-pie/log"
)

type resMovieSearch struct {
	Response     string        `json:"Response"`
	Search       []interface{} `json:"Search"`
	TotalResults int           `json:"total_results"`
}

func SearchMovies(title string, page int) (interface{}, error) {
	log.Log().Infoln("Title : ", title)
	log.Log().Infoln("Page : ", page)

	req, err := gohttp.NewRequest().
		Get(fmt.Sprintf("http://www.omdbapi.com/?s=%s&apikey=c7349bef&page=%d", title, page))

	if err != nil {
		log.Log().Errorln(err)
		return nil, err
	}

	res := resMovieSearch{}
	if err := req.UnmarshalBody(&res); err != nil {
		log.Log().Errorln(err)
		return nil, err
	}

	result := res.Search
	return result, nil
}
