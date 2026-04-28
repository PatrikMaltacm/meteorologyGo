package routes

import (
	"net/http"

	"github.com/PatrikMaltacm/meteorologyGo/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/HealthCheck", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"Status": "Rodando"}) })
	r.POST("/sendData", controllers.SendWheaterData)
}
