package models

type FavouriteMovie struct {
	UserID  int64 `gorm:"primary_key" json:"user_id"`
	MovieID int64 `gorm:"primary_key" json:"movie_id"`
}

func (fm *FavouriteMovie) TableName() string {
	return "favourite_movies"
}
