package main

import (
	"log"

	"github.com/PatrikMaltacm/meteorologyGo/internal/database"
	"github.com/PatrikMaltacm/meteorologyGo/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect("./database.db")
	defer db.Close()

	weatherHandler := handler.NewWheaterHandler(db)

	router := gin.Default()
	handler.InitRoutes(&router.RouterGroup, weatherHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
