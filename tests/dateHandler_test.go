package tests

import (
    "amparo/handlers"
    "amparo/models"
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

var dateLayout = "2006-01-02"

func SetUpRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/important-dates/:userId", handlers.CreateImportantDates)
    r.GET("/important-dates/:userId", handlers.GetImportantDates)
    r.DELETE("/important-dates/:userId", handlers.DeleteImportantDates)
    return r
}

func TestCalculateImportantDates(t *testing.T) {
	deathDate, _ := time.Parse(dateLayout, "2023-10-01")
    dates := handlers.CalculateImportantDates(deathDate)

    assert.Equal(t, "2023-10-04", dates.SeventhDayMass)
    assert.Equal(t, "2023-10-16", dates.DeathRegistration)
    assert.Equal(t, "2023-11-30", dates.InventoryOpening)
    assert.Equal(t, "2023-12-30", dates.DeathPensionRequest)
    assert.Equal(t, "2024-10-01", dates.LifeInsuranceClaim)
}

func TestCreateImportantDates(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    body := map[string]string{
        "dateOfDeath": "2023-10-01",
    }
    jsonValue, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/important-dates/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var response models.ImportantDates
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, 1, response.ID)
    assert.Equal(t, "2023-10-04", response.SeventhDayMass)
    assert.Equal(t, "2023-10-16", response.DeathRegistration)
    assert.Equal(t, "2023-11-30", response.InventoryOpening)
    assert.Equal(t, "2023-12-30", response.DeathPensionRequest)
    assert.Equal(t, "2024-10-01", response.LifeInsuranceClaim)
}

func TestCreateImportantDatesInvalidDate(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    body := map[string]string{
        "dateOfDeath": "data-invalida",
    }
    jsonValue, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/important-dates/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, "Formato de data inválido. Use yyyy-mm-dd", response["error"])
}

func TestCreateImportantDatesFutureDate(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    futureDate := time.Now().AddDate(0, 0, 1).Format(dateLayout)

    body := map[string]string{
        "dateOfDeath": futureDate,
    }
    jsonValue, _ := json.Marshal(body)

    req, _ := http.NewRequest("POST", "/important-dates/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, "Data de óbito não pode ser no futuro", response["error"])
}

func TestCreateImportantDatesAlreadyExists(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    body := map[string]string{
        "dateOfDeath": "2023-10-01",
    }
    jsonValue, _ := json.Marshal(body)
    req, _ := http.NewRequest("POST", "/important-dates/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    req, _ = http.NewRequest("POST", "/important-dates/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusConflict, w.Code)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, "Datas já existem para este usuário", response["error"])
}

func TestGetImportantDates(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    body := map[string]string{
        "dateOfDeath": "2023-10-01",
    }
    jsonValue, _ := json.Marshal(body)
    req, _ := http.NewRequest("POST", "/important-dates/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    req, _ = http.NewRequest("GET", "/important-dates/1", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var response models.ImportantDates
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, 1, response.ID)
    assert.Equal(t, "2023-10-04", response.SeventhDayMass)
    assert.Equal(t, "2023-10-16", response.DeathRegistration)
    assert.Equal(t, "2023-11-30", response.InventoryOpening)
    assert.Equal(t, "2023-12-30", response.DeathPensionRequest)
    assert.Equal(t, "2024-10-01", response.LifeInsuranceClaim)
}

func TestGetImportantDatesNotFound(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    req, _ := http.NewRequest("GET", "/important-dates/999", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, "Registro não encontrado", response["error"])
}

func TestDeleteImportantDates(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    body := map[string]string{
        "dateOfDeath": "2023-10-01",
    }
    jsonValue, _ := json.Marshal(body)
    req, _ := http.NewRequest("POST", "/important-dates/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    req, _ = http.NewRequest("DELETE", "/important-dates/1", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, "Registro deletado com sucesso", response["message"])

    req, _ = http.NewRequest("GET", "/important-dates/1", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteImportantDatesNotFound(t *testing.T) {
    handlers.ResetStore()
    r := SetUpRouter()

    req, _ := http.NewRequest("DELETE", "/important-dates/999", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.Nil(t, err)
    assert.Equal(t, "Registro não encontrado", response["error"])
}
