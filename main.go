package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PatrikMaltacm/meteorologyGo/src/controllers"
	"github.com/PatrikMaltacm/meteorologyGo/src/controllers/routes"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	wheaterController := controllers.NewWheaterController(db)

	fmt.Println("Connected whit database.")

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, wheaterController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
