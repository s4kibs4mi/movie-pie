package api

import (
	"github.com/gin-gonic/gin"
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/repos"
	"net/http"
)

func login(ctx *gin.Context) {
	u, err := repos.NewUserRepo().Register(app.NewScope(app.DB(), ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}
	ctx.JSON(http.StatusCreated, u)
}

func register(ctx *gin.Context) {
	u, err := repos.NewUserRepo().Register(app.NewScope(app.DB(), ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
		return
	}
	ctx.JSON(http.StatusCreated, u)
}

func profile(ctx *gin.Context) {

}
