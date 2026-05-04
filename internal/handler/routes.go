package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, h *WeatherHandler) {
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"Status": "Rodando"}) })

	weather := r.Group("/weather")
	{
		weather.GET("/", h.GetWeatherData)
		weather.GET("/all", h.GetAllWeatherData)
		weather.POST("/", h.SendWeatherData)
	}

	station := r.Group("/station")
	{
		station.GET("/all", h.GetAllStation)
		station.POST("/", h.CreateStation)
	}
}
