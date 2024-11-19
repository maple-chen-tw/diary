package models

import (
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
	return db.Create(diary).Error
}

// UpdateDiary 更新日記
func UpdateDiary(db *gorm.DB, diary *UserDiary) error {
	diary.UpdateAt = time.Now()
	return db.Save(diary).Error
}

// GetDiaryByID 根據 ID 查找日記
func GetDiaryByID(db *gorm.DB, id int) (*UserDiary, error) {
	var diary UserDiary
	err := db.First(&diary, id).Error
	return &diary, err
}

// DeleteDiary 根據 ID 刪除日記
func DeleteDiary(db *gorm.DB, id int) error {
	var diary UserDiary
	err := db.First(&diary, id).Error
	if err != nil {
		return err
	}
	return db.Delete(&diary).Error
}
