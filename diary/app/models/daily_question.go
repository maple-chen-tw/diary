package models

import (
	"gorm.io/gorm"
)

// Create Daily Question
func CreateDailyQuestion(db *gorm.DB, dailyQuestion *DailyQuestion) error {
	return db.Create(dailyQuestion).Error
}

// Get Daily Question By question_id
func GetDailyQuestionByID(db *gorm.DB, id int) (DailyQuestion, error) {
	var dailyQuestion DailyQuestion
	result := db.Where("question_id = ?", id).First(&dailyQuestion)
	return dailyQuestion, result.Error
}

// Get Daily Question By Random
func GetDailyQuestionByRandom(db *gorm.DB) (DailyQuestion, error) {
	var dailyQuestion DailyQuestion
	var result *gorm.DB

	switch db.Dialector.Name() {
	case "mysql":
		result = db.Order("RAND()").First(&dailyQuestion)
	case "sqlite":
		result = db.Order("RANDOM()").First(&dailyQuestion)
	default:
		result = db.Order("RAND()").First(&dailyQuestion) // default to MySQL if dialect is unknown
	}

	return dailyQuestion, result.Error
}
