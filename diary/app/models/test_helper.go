package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, error) {
	// Ensure CGO_ENABLED is set correctly for SQLite
	if os.Getenv("CGO_ENABLED") == "" {
		os.Setenv("CGO_ENABLED", "1")
	}

	// Open an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the UserDiary schema
	err = db.AutoMigrate(&DailyQuestion{}, &UserDiary{}, &User{}, &PublicQuestion{}, &PublicQuestionAnswer{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
