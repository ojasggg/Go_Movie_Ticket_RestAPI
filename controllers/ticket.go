package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ojasggg/movie-theater-gin/models"
)

var tickets []models.Ticket
var ticketIDCounter = 1

func CreateTicket(c *gin.Context) {
	var input models.CreateTicketInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Unauthorized"})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Invalid user ID in token"})
		return
	}

	var movieIndex = -1
	for i, movie := range movies{
		if movie.ID == input.MovieID {
			movieIndex = i
			break
		}
	}

	if movieIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Movie not found"})
		return
	}

	if movies[movieIndex].AvailableSeats < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Not enough seats available"})
		return
	}

	ticket := models.Ticket{
		ID : ticketIDCounter,
		UserID: userID,
		MovieID: input.MovieID,
		Quantity: input.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}	
	tickets = append(tickets, ticket)
	ticketIDCounter ++ 

	movies[movieIndex].AvailableSeats -= input.Quantity

	c.JSON(http.StatusOK, gin.H{"data" : tickets})
}

func GetTickets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data" : tickets})
}

func GetTicketByID(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	for _, ticket := range tickets{
		if ticket.ID == id{
			c.JSON(http.StatusOK, gin.H{"data" : ticket})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Ticket not found"})
}

func UpdateTicket(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	var input models.UpdateTicketInput
	if err:= c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	var ticketIndex = -1
	for i, ticket := range tickets{
		if ticket.ID == id {
			ticketIndex = i
			break
		}
	}

	if ticketIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Ticket not found"})
		return
	}

	oldTicket := tickets[ticketIndex]

	movieIndex := -1
	for i, movie := range movies {
		if movie.ID == oldTicket.MovieID{
			movieIndex = i
			break
		}
	}

	if movieIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Associated Movie not found"})
		return
	}

	seatDiff := input.Quantity - oldTicket.Quantity

	if seatDiff > 0 && movies[movieIndex].AvailableSeats < seatDiff {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Not enough seats available for update"})
		return
	}

	
	tickets[ticketIndex].Quantity = input.Quantity
	tickets[ticketIndex].UpdatedAt = time.Now()
	movies[movieIndex].AvailableSeats -= seatDiff

	c.JSON(http.StatusOK, gin.H{"data":tickets[ticketIndex]})
}

func DeleteTicket(c *gin.Context){
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	for i, ticket := range tickets {
		if ticket.ID == id {
			tickets = append(tickets[:i], tickets[i + 1:]...)
			c.JSON(http.StatusOK, gin.H{"message" : "Ticket deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Ticket not found"})
}