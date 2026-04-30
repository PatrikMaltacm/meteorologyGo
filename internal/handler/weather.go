package handler

import (
	"database/sql"
	"net/http"

	"github.com/PatrikMaltacm/meteorologyGo/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WeatherHandler struct {
	db *sql.DB
}

func NewWeatherHandler(db *sql.DB) *WeatherHandler {
	return &WeatherHandler{
		db: db,
	}
}

func (w *WeatherHandler) SendWeatherData(c *gin.Context) {
	var dataRequest model.WeatherRequest

	if err := c.ShouldBindJSON(&dataRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
			"info":  err.Error(),
		})
		return
	}

	newID := uuid.New().String()
	query := `INSERT INTO weather_data (id, pressure, humidity, temp, lat, long) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := w.db.Exec(query, newID, dataRequest.Pressure, dataRequest.Humidity, dataRequest.Temp, dataRequest.Lat, dataRequest.Long)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"dados": dataRequest})
}

func (w *WeatherHandler) GetAllWeatherData(c *gin.Context) {
	var allData []model.WeatherResponse

	query := `
		SELECT id, pressure, humidity, temp, lat, long, created_at
		FROM weather_data
	`

	rows, err := w.db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar dados."})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var register model.WeatherResponse
		if err := rows.Scan(&register.ID, &register.Pressure, &register.Humidity, &register.Temp, &register.Lat, &register.Long, &register.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao ler registro."})
			return
		}
		allData = append(allData, register)
	}

	c.JSON(http.StatusOK, allData)
}

func (w *WeatherHandler) GetWeatherData(c *gin.Context) {
	var data model.WeatherResponse

	query := `
		SELECT id, pressure, humidity, temp, lat, long, created_at
		FROM weather_data
		ORDER BY created_at DESC
		LIMIT 1
	`

	err := w.db.QueryRow(query).Scan(
		&data.ID,
		&data.Pressure,
		&data.Humidity,
		&data.Temp,
		&data.Lat,
		&data.Long,
		&data.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nenhum dado encontrado."})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (w *WeatherHandler) SetupDatabase(c *gin.Context) {
	query := `
	CREATE TABLE IF NOT EXISTS weather_data (
		id TEXT PRIMARY KEY,
		pressure INTEGER NOT NULL,
		humidity INTEGER NOT NULL,
		temp INTEGER NOT NULL,
		lat REAL NOT NULL,
		long REAL NOT NULL,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := w.db.Exec(query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tabela PostgreSQL pronta"})
}
