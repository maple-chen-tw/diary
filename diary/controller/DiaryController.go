package controller

import (
	models "diary/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var diaries []models.UserDiary

// GetAllDiary
func GetAllDiary(c *gin.Context) {
	c.JSON(http.StatusOK, diaries)
}

// CreateDiary
func CreateDiary(c *gin.Context) {
	var newDiary models.UserDiary
	if err := c.ShouldBindJSON(&newDiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	diaries = append(diaries, newDiary)
	c.JSON(http.StatusCreated, newDiary)
}

// SaveDiary
func SaveDiary(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diary ID"})
		return
	}

	for i, diary := range diaries {
		if diary.ID == id {
			if err := c.ShouldBindJSON(&diaries[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			diaries[i].UpdateAt = time.Now()
			c.JSON(http.StatusOK, diaries[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
}

// UndoDiary
func UndoDiary(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, diary := range diaries {
		if diary.ID == idInt {
			c.JSON(http.StatusOK, diary)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
}

// DeleteDiary
func DeleteDiary(c *gin.Context) {
	idStr := c.Param("question_id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for i, diary := range diaries {
		if diary.ID == idInt {
			diaries = append(diaries[:i], diaries[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Diary not found"})
}

// DeleteAllDiary
func DeleteAllDiary(c *gin.Context) {
	diaries = []models.UserDiary{}
	c.Status(http.StatusNoContent)
}
