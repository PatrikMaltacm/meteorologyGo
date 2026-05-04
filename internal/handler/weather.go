package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/PatrikMaltacm/meteorologyGo/internal/model"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WeatherHandler struct {
	db    *sql.DB
	cache *bigcache.BigCache
}

func NewWeatherHandler(db *sql.DB, cache *bigcache.BigCache) *WeatherHandler {
	return &WeatherHandler{
		db:    db,
		cache: cache,
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
	now := time.Now()

	query := `INSERT INTO weather_data (id, station_id, pressure, humidity, temp, lat, long, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := w.db.Exec(query, newID, dataRequest.StationID, dataRequest.Pressure, dataRequest.Humidity, dataRequest.Temp, dataRequest.Lat, dataRequest.Long, now)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cacheData := model.WeatherResponse{
		ID:        newID,
		StationID: dataRequest.StationID,
		Pressure:  dataRequest.Pressure,
		Humidity:  dataRequest.Humidity,
		Temp:      dataRequest.Temp,
		Lat:       dataRequest.Lat,
		Long:      dataRequest.Long,
		CreatedAt: now,
	}

	if jsonData, err := json.Marshal(cacheData); err == nil {
		w.cache.Set("latest_weather", jsonData)
	}

	c.JSON(http.StatusAccepted, gin.H{"dados": cacheData})
}

func (w *WeatherHandler) GetAllWeatherData(c *gin.Context) {
	var allData []model.WeatherResponse

	query := `
		SELECT id, pressure, humidity, temp, station_id, lat, long, created_at
		FROM weather_data LIMIT 100
	`

	rows, err := w.db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar dados."})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var register model.WeatherResponse
		if err := rows.Scan(&register.ID, &register.Pressure, &register.Humidity, &register.Temp, &register.StationID, &register.Lat, &register.Long, &register.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao ler registro."})
			return
		}
		allData = append(allData, register)
	}

	c.JSON(http.StatusOK, allData)
}

func (w *WeatherHandler) GetWeatherData(c *gin.Context) {
	if entry, err := w.cache.Get("latest_weather"); err == nil {
		var cachedData model.WeatherResponse
		if err := json.Unmarshal(entry, &cachedData); err == nil {
			c.JSON(http.StatusOK, cachedData)
			return
		}
	}

	var data model.WeatherResponse
	query := `
        SELECT id, pressure, humidity, temp, station_id, lat, long, created_at 
        FROM weather_data 
        ORDER BY created_at DESC 
        LIMIT 1
    `

	err := w.db.QueryRow(query).Scan(
		&data.ID,
		&data.Pressure,
		&data.Humidity,
		&data.Temp,
		&data.StationID,
		&data.Lat,
		&data.Long,
		&data.CreatedAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nenhum dado encontrado."})
		return
	}

	if jsonData, err := json.Marshal(data); err == nil {
		w.cache.Set("latest_weather", jsonData)
	}

	c.JSON(http.StatusOK, data)
}
