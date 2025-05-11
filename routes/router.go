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
		movieRoutes := api.Group("/movies")
		{
			movieRoutes.POST("", controllers.CreateMovie)
			movieRoutes.GET("/:id", controllers.GetMovieByID)
			movieRoutes.GET("", controllers.GetMovies)
			movieRoutes.PUT("/:id", controllers.UpdateMovie)
			movieRoutes.DELETE("/:id",controllers.DeleteMovie)
		}

		ticketRoutes := api.Group("/tickets")
		{
			ticketRoutes.POST("", controllers.CreateTicket)
			ticketRoutes.GET("/:id", controllers.GetTicketByID)
			ticketRoutes.GET("", controllers.GetTickets)
			ticketRoutes.PUT("/:id", controllers.UpdateTicket)
			ticketRoutes.DELETE("/:id", controllers.DeleteTicket)
		}

		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", controllers.Register)
			authRoutes.POST("/login", controllers.Login)

		}
	}
}