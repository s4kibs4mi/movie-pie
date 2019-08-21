package models

type FavouriteMovie struct {
	UserID  int64 `json:"user_id"`
	MovieID int64 `json:"movie_id"`
}

func (fm *FavouriteMovie) TableName() string {
	return "favourite_movies"
}
