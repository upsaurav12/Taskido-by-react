package main

import (
	"log"
	"taskido/backend/database"
	"taskido/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.Connect()

	// Setup routes
	r := routes.Setuproutes()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Add logging middleware
	r.Use(func(c *gin.Context) {
		log.Println("Request Headers:", c.Request.Header)
		c.Next()
	})

	// Add recovery middleware
	r.Use(gin.Recovery())

	// Start the server
	port := ":8000"
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
