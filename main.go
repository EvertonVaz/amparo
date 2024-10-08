package main

import (
	"amparo/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define os endpoints da API
	router.POST("/important-dates/:userId", handlers.CreateImportantDates)
	router.GET("/important-dates/:userId", handlers.GetImportantDates)
	router.DELETE("/important-dates/:userId", handlers.DeleteImportantDates)

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}
