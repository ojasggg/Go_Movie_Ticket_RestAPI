package models

import "time"

type Ticket struct {
	ID int `json:"id"`
	MovieID int `json:"movieId"`
	UserID int `json:"userId"`
	Quantity int `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTicketInput struct {
	MovieID int `json:"movieId" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type UpdateTicketInput struct {
	Quantity int `json:"quantity" binding:"required"`
}