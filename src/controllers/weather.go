package controllers

import (
	"database/sql"
	"net/http"

	"github.com/PatrikMaltacm/meteorologyGo/src/models/request"
	"github.com/gin-gonic/gin"
)

type WheaterController struct {
	db *sql.DB
}

func NewWheaterController(db *sql.DB) *WheaterController {
	return &WheaterController{
		db: db,
	}
}

func (w *WheaterController) SendWheaterData(c *gin.Context) {
	var dataRequest request.WheatherRequest

	c.ShouldBindJSON(&dataRequest)

	c.JSON(http.StatusAccepted, gin.H{"message:": "Dados Enviados ao servidor"})
}
