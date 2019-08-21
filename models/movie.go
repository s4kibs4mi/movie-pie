package models

import "time"

type Movie struct {
	ID        int64     `json:"id" sql:"primary_key;auto_increment"`
	Title     string    `json:"title" sql:"index"`
	Year      int       `json:"year"`
	IMDBID    string    `json:"imdbid" sql:"index;unique"`
	Poster    string    `json:"poster"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *Movie) TableName() string {
	return "movies"
}
