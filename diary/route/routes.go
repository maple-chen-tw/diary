package routes

import (
	"diary/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/diary", controller.GetAllDiary)
	router.POST("/api/diary", controller.CreateDiary)
	router.PUT("/api/diary/:id", controller.SaveDiary)
	router.PUT("/api/undoDiary/:id", controller.UndoDiary)
	router.DELETE("/api/deleteDiary/:id", controller.DeleteDiary)

	return router
}
