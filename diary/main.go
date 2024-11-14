package main

import (
	"diary/database"
	"diary/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := database.InitializeDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // 前端的地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允許的方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允許的標頭
		AllowCredentials: true,
	}))

	routes.SetupRoutes(router, db)

	// Start the server on port 8080
	log.Println("Server is running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
