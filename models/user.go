package models

import "time"

type User struct {
	ID int `json:"id"`
	Username string `json:"username" binding:"required"`
	Email string `josn:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Email string `josn:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginUser struct {
	Email string `josn:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}