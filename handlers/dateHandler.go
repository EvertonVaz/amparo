package handlers

import (
    "amparo/models"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

var importantDatesStore = make(map[string]models.ImportantDates)
var idCounter = 1

func ResetStore() {
    importantDatesStore = make(map[string]models.ImportantDates)
    idCounter = 1
}

func CalculateImportantDates(deathDate time.Time) models.ImportantDates {
    layout := "2006-01-02"
    dates := models.ImportantDates{
        SeventhDayMass:      deathDate.AddDate(0, 0, 3).Format(layout),
        DeathRegistration:   deathDate.AddDate(0, 0, 15).Format(layout),
        InventoryOpening:    deathDate.AddDate(0, 0, 60).Format(layout),
        DeathPensionRequest: deathDate.AddDate(0, 0, 90).Format(layout),
        LifeInsuranceClaim:  deathDate.AddDate(1, 0, 0).Format(layout),
    }
    return dates
}

func CreateImportantDates(c *gin.Context) {
    userId := c.Param("userId")

    if _, exists := importantDatesStore[userId]; exists {
        c.JSON(http.StatusConflict, gin.H{"error": "Datas já existem para este usuário"})
        return
    }

    var request struct {
        DateOfDeath string `json:"dateOfDeath"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválido"})
        return
    }

    layout := "2006-01-02"
    deathDate, err := time.Parse(layout, request.DateOfDeath)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use yyyy-mm-dd"})
        return
    }

    if deathDate.After(time.Now()) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data de óbito não pode ser no futuro"})
        return
    }

    dates := CalculateImportantDates(deathDate)
    dates.ID = idCounter
    idCounter++

    importantDatesStore[userId] = dates

    c.JSON(http.StatusOK, dates)
}

func GetImportantDates(c *gin.Context) {
    userId := c.Param("userId")
    dates, exists := importantDatesStore[userId]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
        return
    }

    c.JSON(http.StatusOK, dates)
}

func DeleteImportantDates(c *gin.Context) {
    userId := c.Param("userId")
    if _, exists := importantDatesStore[userId]; !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
        return
    }

    delete(importantDatesStore, userId)
    c.JSON(http.StatusOK, gin.H{"message": "Registro deletado com sucesso"})
}
