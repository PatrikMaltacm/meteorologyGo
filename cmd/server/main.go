package main

import (
	"log"
	"os"

	"github.com/PatrikMaltacm/meteorologyGo/internal/database"
	"github.com/PatrikMaltacm/meteorologyGo/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "./database.db"
	}

	db := database.Connect(dsn)
	defer db.Close()

	weatherHandler := handler.NewWeatherHandler(db)

	router := gin.Default()
	handler.InitRoutes(&router.RouterGroup, weatherHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
