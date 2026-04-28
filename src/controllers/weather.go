package controllers

import (
	"net/http"

	"github.com/PatrikMaltacm/meteorologyGo/src/models/request"
	"github.com/gin-gonic/gin"
)

func SendWheaterData(c *gin.Context) {
	var dataRequest request.WheatherRequest

	c.ShouldBindJSON(&dataRequest)

	c.JSON(http.StatusAccepted, gin.H{"message:": "Dados Enviados ao servidor"})
}
