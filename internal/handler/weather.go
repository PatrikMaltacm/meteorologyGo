package handler

import (
	"database/sql"
	"net/http"

	"github.com/PatrikMaltacm/meteorologyGo/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WheaterHandler struct {
	db *sql.DB
}

func NewWheaterHandler(db *sql.DB) *WheaterHandler {
	return &WheaterHandler{
		db: db,
	}
}

func (w *WheaterHandler) SendWheaterData(c *gin.Context) {
	var dataRequest model.WheatherRequest

	if err := c.ShouldBindJSON(&dataRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
			"info":  err.Error(),
		})
		return
	}

	newID := uuid.New().String()
	query := `INSERT INTO weather_data (id, pressure, humidity, temp) VALUES (?, ?, ?, ?)`

	_, err := w.db.Exec(query, &newID, &dataRequest.Pressure, &dataRequest.Humidity, &dataRequest.Temp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"dados:": dataRequest})
}

func (w *WheaterHandler) GetWheaterData(c *gin.Context) {
	var data model.WheaterResponse

	query := `
        SELECT 
            id, pressure, humidity, temp, created_at
        FROM weather_data 
        ORDER BY created_at DESC 
        LIMIT 1
    `

	err := w.db.QueryRow(query).Scan(
		&data.ID,
		&data.Pressure,
		&data.Humidity,
		&data.Temp,
		&data.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message:": "Nenhum dado encontrado."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Dados:": data})
}

func (w *WheaterHandler) SetupDatabase(c *gin.Context) {
	query := `
    CREATE TABLE IF NOT EXISTS weather_data (
        id TEXT PRIMARY KEY,
        pressure INTEGER NOT NULL,
        humidity INTEGER NOT NULL,
        temp INTEGER NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	if _, err := w.db.Exec(query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tabela SQLite pronta"})
}
