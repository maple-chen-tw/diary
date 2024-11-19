package controllers

import (
	"diary/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateDailyQuestion 創建一條新的問題記錄
func CreateDailyQuestion(c *gin.Context, db *gorm.DB) {
	var newQuestion models.DailyQuestion
	// 綁定請求中的 JSON 資料到 newQuestion
	if err := c.ShouldBindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 呼叫model layer創建問題
	if err := models.CreateDailyQuestion(db, &newQuestion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create daily question"})
		return
	}

	// 返回創建成功的問題記錄
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Daily question created successfully",
		"question": newQuestion,
	})
}

// GetDailyQuestionByID 根據問題 ID 獲取問題記錄
func GetDailyQuestionByID(c *gin.Context, db *gorm.DB) {
	// 從 URL 參數中提取問題 ID
	idStr := c.Param("id")

	// 將 string 類型的 id 轉換為 int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// 如果轉換失敗，返回錯誤訊息
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// 根據 ID 查找問題
	question, err := models.GetDailyQuestionByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// 返回問題記錄
	c.JSON(http.StatusOK, gin.H{
		"question": question,
	})
}

// GetDailyQuestionByRandom 隨機獲取一條問題記錄
func GetDailyQuestionByRandom(c *gin.Context, db *gorm.DB) {
	// 獲取隨機問題
	question, err := models.GetDailyQuestionByRandom(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch random question"})
		return
	}

	// 返回隨機問題記錄
	c.JSON(http.StatusOK, gin.H{
		"question": question,
	})
}
