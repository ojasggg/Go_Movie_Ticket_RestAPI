package controllers

import (
	"net/http"
	"strconv"
	"time"

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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func UpdateMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	} 

	var input models.CreateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies[i].Title = input.Title
			movies[i].Description = input.Description
			movies[i].ShowTime = input.ShowTime
			movies[i].AvailableSeats = input.AvailableSeats
			movies[i].TotalSeats = input.TotalSeats
			movies[i].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, gin.H{"data" : movies[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Movie not found"})
}

func DeleteMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	for i, movie := range movies {
		if movie.ID == id{
			movies = append(movies[:i], movies[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message" : "Movie deleted"})
			return 
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}