package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mta-hosting-optimizer/handlers"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/hostnames", handlers.GetHostnames)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...\n", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
