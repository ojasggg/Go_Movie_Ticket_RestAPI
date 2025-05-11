package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ojasggg/movie-theater-gin/models"
	"github.com/ojasggg/movie-theater-gin/utils"
	"golang.org/x/crypto/bcrypt"
)

var users []models.User
var userIDCounter = 1

func init() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash default admin password")
	}

	adminUser := models.User{
		ID : userIDCounter,
		Username: "admin",
		Email: "admin@example.com",
		Role: "admin",
		Password: string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, adminUser)
	userIDCounter ++
}

func Register (c *gin.Context) {
	var input models.RegisterUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	user := models.User{
		ID : userIDCounter,
		Username: input.Username,
		Email: input.Email,
		Password: string(hashedPassword),
		Role: "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, user)
	userIDCounter ++

	c.JSON(http.StatusOK, gin.H{
		"data" : gin.H{
			"id" : user.ID,
			"username" : user.Username,
 			"email" : user.Email,
		},
 	})
}

func Login(c *gin.Context){
	var input models.LoginUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	for _, user := range users{
		if user.Email == input.Email{
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil{
				c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid Credentials"})
				return
			}

			token, err := utils.GenerateJWT(user.ID, user.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"token" : token,
				"user" : gin.H{
					"id" : user.ID,
					"username": user.Username,
					"email" : user.Email,
				},
			})
			return
		}	
	}

	c.JSON(http.StatusNotFound, gin.H{"error" : "Invalid Credentials"})
}