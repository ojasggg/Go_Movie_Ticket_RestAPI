package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ojasggg/movie-theater-gin/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/movies", controllers.CreateMovie)
		api.GET("/movies:id", controllers.GetMovieByID)
		api.GET("/movies", controllers.GetMovies)

	}
}