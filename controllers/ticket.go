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

	ticket := models.Ticket{
		ID : ticketIDCounter,
		UserID: input.UserID,
		MovieID: input.MovieID,
		Quantity: input.Quantity,
		CreatedAt: time.Now(),

	}	
	tickets = append(tickets, ticket)
	ticketIDCounter ++ 

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

	for i, ticket := range tickets{
		if ticket.ID == id {
			tickets[i].Quantity = input.Quantity

			c.JSON(http.StatusOK, gin.H{"data" : tickets[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
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