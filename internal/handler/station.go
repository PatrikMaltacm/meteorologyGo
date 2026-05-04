package handler

import (
	"net/http"

	"github.com/PatrikMaltacm/meteorologyGo/internal/model"
	"github.com/gin-gonic/gin"
)

func (w *WeatherHandler) GetAllStation(c *gin.Context) {
	var allData []model.StationResponse

	query := `
		SELECT id, lat, long, created_at
		FROM stations
	`

	rows, err := w.db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar dados."})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var register model.StationResponse
		if err := rows.Scan(&register.ID, &register.Lat, &register.Long, &register.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao ler registro."})
			return
		}
		allData = append(allData, register)
	}

	c.JSON(http.StatusOK, allData)
}

func (w *WeatherHandler) CreateStation(c *gin.Context) {
	var dataRequest model.CreateStationResponse

	if err := c.ShouldBindJSON(&dataRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
			"info":  err.Error(),
		})
		return
	}

	query := `INSERT INTO stations (lat, long) VALUES ($1, $2)`

	_, err := w.db.Exec(query, dataRequest.Lat, dataRequest.Long)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"dados": dataRequest})
}
