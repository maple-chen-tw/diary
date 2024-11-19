package controllers

import (
	models "diary/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllDiary 根據用戶 ID 獲取所有日記
func GetAllDiary(c *gin.Context, db *gorm.DB) {
	// 從 JWT token 中提取用戶的 user_id
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	// 強制轉換為 int 型別
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// 根據 user_id 獲取該用戶的所有日記
	diaries, err := models.GetAllDiariesByUserID(db, userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, diaries)
}

// CreateDiary 創建新的日記
func CreateDiary(c *gin.Context, db *gorm.DB) {
	var newDiary models.UserDiary
	if err := c.ShouldBindJSON(&newDiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 從 JWT token 中提取用戶的 user_id
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	// 強制轉換為 int 型別
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// 設置用戶 ID
	newDiary.UserID = userIDInt
	newDiary.CreateAt = time.Now()
	newDiary.UpdateAt = time.Now()

	if err := models.CreateDiary(db, &newDiary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create diary"})
		return
	}

	c.JSON(http.StatusCreated, newDiary)
}

// SaveDiary
func SaveDiary(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diary ID"})
		return
	}

	diary, err := models.GetDiaryByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
		return
	}

	// 檢查日記是否屬於該用戶
	userID, exists := c.Get("user_id")
	if !exists || diary.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this diary"})
		return
	}

	if err := c.ShouldBindJSON(&diary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateDiary(db, diary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update diary"})
		return
	}

	c.JSON(http.StatusOK, diary)
}

// UndoDiary 根據 ID 獲取日記
func UndoDiary(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	diary, err := models.GetDiaryByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
		return
	}

	// 檢查日記是否屬於該用戶
	userID, exists := c.Get("user_id")
	if !exists || diary.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to view this diary"})
		return
	}

	c.JSON(http.StatusOK, diary)
}

// DeleteDiary 刪除日記
func DeleteDiary(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// 根據日記 ID 獲取日記
	diary, err := models.GetDiaryByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
		return
	}

	// 檢查日記是否屬於該用戶
	userID, exists := c.Get("user_id")
	if !exists || diary.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this diary"})
		return
	}

	if err := models.DeleteDiary(db, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete diary"})
		return
	}

	c.Status(http.StatusNoContent)
}
