package api

import (
	"github.com/gin-gonic/gin"
	"github.com/s4kibs4mi/movie-pie/app"
	"github.com/s4kibs4mi/movie-pie/data"
	"net/http"
	"strings"
)

func auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authPayload := ctx.GetHeader("Authorization")
		values := strings.Split(authPayload, "Bearer ")

		if len(values) < 1 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized user",
			})
			return
		}

		token := values[1]

		userDao := data.NewUserDao()
		if err := userDao.ValidateToken(app.NewScope(app.DB(), nil), token); err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized user",
			})
			return
		}

		ctx.Next()
	}
}
