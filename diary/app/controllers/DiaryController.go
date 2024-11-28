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

	diaries, err := models.GetAllDiariesByUserID(db, userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, diaries)
}

// Create New Diary
func CreateDiary(c *gin.Context, db *gorm.DB) {
	var newDiary models.UserDiary
	if err := c.ShouldBindJSON(&newDiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user_id from JWT token
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

	// set up user_id
	newDiary.UserID = userIDInt

	if err := models.CreateDiary(db, &newDiary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create diary"})
		return
	}

	if newDiary.DiaryID == 0 {
		// Ensure the diary_id is correctly populated by fetching the last record
		db.Last(&newDiary)
		log.Printf("Manual ID retrieval: New diary created with ID: %d", newDiary.DiaryID)
	}
	log.Printf("New diary created with ID: %d", newDiary.DiaryID)
	c.JSON(http.StatusCreated, newDiary)
}

// SaveDiary
func SaveDiary(c *gin.Context, db *gorm.DB) {

	idStr := c.Param("diary_id")
	diaryID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diary ID"})
		return
	}

	userDiary, err := models.GetUserDiary(db, diaryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
		return
	}
	// check if user_id have authorization of the diary
	userID, exists := c.Get("user_id")
	if !exists || userDiary.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to update this diary"})
		return
	}

	if err := c.ShouldBindJSON(&userDiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDiary.UpdateAt = time.Now()

	// update diary data
	updatedDiary, err := models.UpdateUserDiary(db, userDiary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update diary"})
		return
	}
	c.JSON(http.StatusOK, updatedDiary)
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
