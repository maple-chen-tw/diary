package models

import (
	"gorm.io/gorm"
)

// CreateDailyQuestion 創建一條新的問題記錄
func CreateDailyQuestion(db *gorm.DB, dailyQuestion *DailyQuestion) error {
	return db.Create(dailyQuestion).Error
}

// GetDailyQuestionByID 根據問題 ID 獲取問題記錄
func GetDailyQuestionByID(db *gorm.DB, id int) (DailyQuestion, error) {
	var dailyQuestion DailyQuestion
	result := db.Where("id = ?", id).First(&dailyQuestion)
	return dailyQuestion, result.Error
}

// GetDailyQuestionByRandom 根據隨機的問題 ID 獲取問題記錄
func GetDailyQuestionByRandom(db *gorm.DB) (DailyQuestion, error) {
	var dailyQuestion DailyQuestion
	result := db.Order("RAND()").First(&dailyQuestion)
	return dailyQuestion, result.Error
}

/*
// GetDailyQuestions 獲取所有問題記錄
func GetDailyQuestions(db *gorm.DB) ([]DailyQuestion, error) {
	var dailyQuestions []DailyQuestion
	result := db.Find(&dailyQuestions)
	return dailyQuestions, result.Error
}

// UpdateDailyQuestion 更新問題記錄
func UpdateDailyQuestion(db *gorm.DB, dailyQuestion *DailyQuestion) error {
	return db.Save(dailyQuestion).Error
}

// DeleteDailyQuestion 刪除問題記錄
func DeleteDailyQuestion(db *gorm.DB, id int) error {
	return db.Delete(&DailyQuestion{}, id).Error
}

*/
