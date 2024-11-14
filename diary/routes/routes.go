package routes

import (
	"diary/controllers"
	"diary/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes sets up the application routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/register", func(c *gin.Context) {
		controllers.Register(c, db)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	// 需要 JWT 驗證的路由
	protected := router.Group("/protected")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(200, gin.H{"message": "Welcome to your profile", "username": username})
		})
	}

	//router.GET("/api/diary", controllers.GetAllDiary)
	//router.POST("/api/diary", controller.CreateDiary)
	//router.PUT("/api/diary/:id", controller.SaveDiary)
	//router.PUT("/api/undoDiary/:id", middleware.UndoDiary)
	//router.DELETE("/api/deleteDiary/:id", controller.DeleteDiary)
}
