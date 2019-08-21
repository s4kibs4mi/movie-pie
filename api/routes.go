package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var router = gin.Default()

// Router returns the api router
func Router() http.Handler {
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(func(ctx *gin.Context) {
		defer func() {
			if rvr := recover(); rvr != nil {
				resp := Response{}
				resp.Title = "Something went wrong"
				resp.Errors = rvr.(error)
				resp.Status = http.StatusInternalServerError
				resp.ServerJSON(ctx.Writer)
				return
			}
		}()
	})

	router.GET("/", func(ctx *gin.Context) {
		resp := Response{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"name": "movie-pie",
			},
		}
		resp.ServerJSON(ctx.Writer)
	})

	registerRoutes()

	return router
}

func registerRoutes() {
	v1 := router.Group("/v1")

	v1.POST("/login", login)
	v1.POST("/register", register)

	v1.Use(auth()).GET("/profile", profile)

	movies := v1.Group("movies")
	movies.Use(auth())
	{
		movies.GET("/search", searchMovie)
		movies.GET("/favourite", favouriteMovie)
	}
}
