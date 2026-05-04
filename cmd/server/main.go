package main

import (
	"log"
	"os"

	"github.com/PatrikMaltacm/meteorologyGo/internal/cache"
	"github.com/PatrikMaltacm/meteorologyGo/internal/database"
	"github.com/PatrikMaltacm/meteorologyGo/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado")
	}

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	c, err := cache.NewBigCache()
	if err != nil {
		log.Fatal("Erro ao iniciar cache:", err)
	}

	db := database.Connect(dsn)
	defer db.Close()

	weatherHandler := handler.NewWeatherHandler(db, c)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	handler.InitRoutes(v1, weatherHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
