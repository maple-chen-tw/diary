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
	protected := router.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(200, gin.H{"message": "Welcome to your profile", "username": username})
		})

		// Diary related routes
		protected.GET("/diary", func(c *gin.Context) {
			controllers.GetAllDiary(c, db)
		})

		protected.POST("/diary", func(c *gin.Context) {
			controllers.CreateDiary(c, db)
		})

		protected.PUT("/diary/:id", func(c *gin.Context) {
			controllers.SaveDiary(c, db)
		})

		protected.GET("/diary/:id", func(c *gin.Context) {
			controllers.GetDiary(c, db)
		})

		protected.DELETE("/diary/:id", func(c *gin.Context) {
			controllers.DeleteDiary(c, db)
		})
		// Daily question related routes
		protected.POST("/question", func(c *gin.Context) {
			controllers.CreateDailyQuestion(c, db)
		})

		protected.GET("/question/:id", func(c *gin.Context) {
			controllers.GetDailyQuestionByID(c, db)
		})

		protected.GET("/question/random", func(c *gin.Context) {
			controllers.GetDailyQuestionByRandom(c, db)
		})

	}
}
