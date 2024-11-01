package routes

import (
	"diary/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/register", func(c *gin.Context) {
		controllers.Register(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	//router.GET("/api/diary", controller.GetAllDiary)
	//router.POST("/api/diary", controller.CreateDiary)
	//router.PUT("/api/diary/:id", controller.SaveDiary)
	//router.PUT("/api/undoDiary/:id", controller.UndoDiary)
	//router.DELETE("/api/deleteDiary/:id", controller.DeleteDiary)

	//return router
}
