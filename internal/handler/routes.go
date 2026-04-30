package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, h *WeatherHandler) {
	r.GET("/HealthCheck", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"Status": "Rodando"}) })
	r.GET("/getData", h.GetWeatherData)
	r.GET("/getAllData", h.GetAllWeatherData)
	r.POST("/sendData", h.SendWeatherData)
	r.POST("/setup", h.SetupDatabase)
}
