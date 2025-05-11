package models

import "time"

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `josn:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Email string `josn:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email string `josn:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}