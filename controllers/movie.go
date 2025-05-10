package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ojasggg/movie-theater-gin/models"
)

var movies []models.Movie
var movieIdCounter = 1

func CreateMovie(c *gin.Context){
	var input models.CreateMovieInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	movie := models.Movie{
		ID: movieIdCounter,
		Title: input.Title,
		Description: input.Description,
		ShowTime: input.ShowTime,
		TotalSeats: input.TotalSeats,
		AvailableSeats: input.AvailableSeats,
	}

	movies = append(movies, movie)
	movieIdCounter ++ 

	c.JSON(http.StatusCreated, gin.H{"data" : movie})
}

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

func GetMovieByID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid movie ID"})
		return
	}

	for _, movie := range movies{
		if movie.ID == id{
			c.JSON(http.StatusOK, gin.H{"data" : movie})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Movie not found"})
}