package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ojasggg/movie-theater-gin/routes"
)

func main() {
	log.Println("🚀 Server has (re)started – memory cleared")
	r := gin.Default()

	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}