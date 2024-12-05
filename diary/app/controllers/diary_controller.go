package controllers

import (
	models "diary/app/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get all diaries by user_id
func GetAllDiary(c *gin.Context, db *gorm.DB) {

	// get user_id from JWT token
	rawID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	id, ok := rawID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	diaries, err := models.GetAllDiariesByUserID(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, diaries)
}

// Create New Diary
func CreateDiary(c *gin.Context, db *gorm.DB) {
	var d models.UserDiary
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user_id from JWT token
	rawID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	id, ok := rawID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// set up user_id
	d.UserID = id
	d.CreateAt = time.Now()
	d.UpdateAt = time.Now()
	if err := models.CreateDiary(db, &d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create diary"})
		return
	}
	log.Printf("New diary created with ID: %d", d.DiaryID)
	c.JSON(http.StatusCreated, d)
}

// SaveDiary
func SaveDiary(c *gin.Context, db *gorm.DB) {

	rawDiaryID := c.Param("diary_id")
	diaryID, err := strconv.Atoi(rawDiaryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diary ID"})
		return
	}

	rawDiary, err := models.GetUserDiary(db, diaryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
		return
	}
	// check if user_id have authorization of the diary
	userID, exists := c.Get("user_id")
	if !exists || rawDiary.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this diary"})
		return
	}

	if err := c.ShouldBindJSON(&rawDiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rawDiary.UpdateAt = time.Now()

	// update diary data
	d, err := models.UpdateUserDiary(db, rawDiary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update diary"})
		return
	}
	c.JSON(http.StatusOK, d)
}

// GetDiary
func GetDiary(c *gin.Context, db *gorm.DB) {
	// Step 1: Get the userID from the JWT token
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Step 2: Get the diary ID from the URL parameter
	idStr := c.Param("diary_id")
	diaryID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diary ID"})
		return
	}

	// Step 3: Fetch the diary from the database
	userDiary, err := models.GetUserDiary(db, diaryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
		return
	}

	// Step 4: Check if the diary belongs to the user from the JWT
	if userDiary.UserID != userIDInt {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this diary"})
		return
	}

	c.JSON(http.StatusOK, userDiary)
}

// DeleteDiary
func DeleteDiary(c *gin.Context, db *gorm.DB) {
	// Get diary ID from the URL parameters
	id := c.Param("diary_id") // ID is passed as a URL parameter, e.g. /diaries/:diary_id
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diary ID is required"})
		return
	}

	// Convert the ID to an integer
	diaryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid diary ID"})
		return
	}

	err = models.DeleteDiary(db, diaryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "diary not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete diary", "message": err.Error()})
		}
		return
	}

	// Respond with success message (204 No Content is ideal for deletions)
	c.JSON(http.StatusNoContent, nil)
}
