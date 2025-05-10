package models

type Movie struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ShowTime string `json:"showTime"`
	TotalSeats int `json:"totalSeats"`
	AvailableSeats int `json:"availableSeats"`
}

type CreateMovieInput struct{
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ShowTime string `json:"showTime" binding:"required"`
	TotalSeats int `json:"totalSeats" binding:"required"`
	AvailableSeats int `json:"availableSeats" binding:"required"`
}