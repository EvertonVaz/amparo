package main

import (
	"amparo/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/important-dates/:userId", handlers.CreateImportantDates)
	router.GET("/important-dates/:userId", handlers.GetImportantDates)
	router.DELETE("/important-dates/:userId", handlers.DeleteImportantDates)

	router.Run(":8080")
}
