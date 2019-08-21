package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/repos"
	"net/http"
)

func login(ctx *gin.Context) {
	u, err := repos.NewUserRepo().Login(app.NewScope(app.DB(), ctx))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "email or password mismatch",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func register(ctx *gin.Context) {
	u, err := repos.NewUserRepo().Register(app.NewScope(app.DB(), ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, u)
}

func profile(ctx *gin.Context) {

}
