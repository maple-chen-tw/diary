package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// GetAllDiariesByUserID 根據用戶 ID 獲取所有日記
func GetAllDiariesByUserID(db *gorm.DB, userID int) ([]UserDiary, error) {
	var diaries []UserDiary
	err := db.Where("user_id = ?", userID).Find(&diaries).Error
	return diaries, err
}

// CreateDiary 創建新的日記
func CreateDiary(db *gorm.DB, diary *UserDiary) error {
	diary.CreateAt = time.Now()
	diary.UpdateAt = time.Now()
	if err := db.Create(diary).Error; err != nil {
		return err
	}
	if diary.DiaryID == 0 {
		if err := db.Last(diary).Error; err != nil {
			return err
		}
		log.Printf("Manual ID retrieval: New diary created with ID: %d", diary.DiaryID)
	}

	return nil
}

func UpdateUserDiary(db *gorm.DB, diary *UserDiary) (*UserDiary, error) {

	// Set the updated timestamp
	diary.UpdateAt = time.Now()

	// Perform the update operation
	if err := db.Model(&UserDiary{}).
		Where("diary_id = ?", diary.DiaryID).
		Updates(map[string]interface{}{
			"question_id": diary.QuestionID,
			"question":    diary.Question,
			"content":     diary.Content,
			"update_at":   diary.UpdateAt,
		}).Error; err != nil {
		return nil, err
	}

	// Return the updated diary
	return diary, nil
}

// Get a UserDiary by diaryID
func GetUserDiary(db *gorm.DB, diaryID int) (*UserDiary, error) {
	var diary UserDiary
	if err := db.Where("diary_id = ?", diaryID).First(&diary).Error; err != nil {
		return nil, err
	}
	return &diary, nil
}

// Delete a UserDiary by diaryID
func DeleteDiary(db *gorm.DB, id int) error {
	var diary UserDiary
	err := db.First(&diary, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("diary not found")
		}
		return err
	}
	err = db.Delete(&diary, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete diary: %v", err)
	}

	return nil
}
