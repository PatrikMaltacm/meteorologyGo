package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, h *WheaterHandler) {
	r.GET("/HealthCheck", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"Status": "Rodando"}) })
	r.GET("/getData", h.GetWheaterData)
	r.POST("/sendData", h.SendWheaterData)
	r.POST("/setup", h.SetupDatabase)
}
