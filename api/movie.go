package api

import (
	"github.com/gin-gonic/gin"
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/repos"
	"net/http"
)

func searchMovie(ctx *gin.Context) {
	movies, err := repos.NewMovieRepo().Search(app.NewScope(app.DB(), ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": movies,
	})
}

func favouriteMovie(ctx *gin.Context) {

}
