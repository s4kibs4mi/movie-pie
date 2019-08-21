package models

type Movie struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	IMDBID string `json:"imdbid"`
	Poster string `json:"poster"`
}

func (m *Movie) TableName() string {
	return "movies"
}
