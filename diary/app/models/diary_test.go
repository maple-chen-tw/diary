package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetAllDiariesByUserID
func TestGetAllDiariesByUserID(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Insert test data
	userID := 1
	diary1 := &UserDiary{UserID: userID, Question: "Question 1", Content: "Content 1"}
	diary2 := &UserDiary{UserID: userID, Question: "Question 2", Content: "Content 2"}
	db.Create(diary1)
	db.Create(diary2)

	// Test GetAllDiariesByUserID
	diaries, err := GetAllDiariesByUserID(db, userID)
	assert.Nil(t, err)
	assert.Len(t, diaries, 2)
	assert.Equal(t, diaries[0].UserID, userID)
}

// Test CreateDiary
func TestCreateDiary(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Create a new diary
	diary := &UserDiary{UserID: 1, Question: "New Question", Content: "New Content"}
	err = CreateDiary(db, diary)
	assert.Nil(t, err)
	assert.NotZero(t, diary.DiaryID) // Ensure diary ID is set after creation
}

// Test UpdateUserDiary
func TestUpdateUserDiary(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Insert a test diary
	diary := &UserDiary{UserID: 1, Question: "Initial Question", Content: "Initial Content"}
	err = CreateDiary(db, diary)
	assert.Nil(t, err)

	// Update the diary
	diary.Question = "Updated Question"
	diary.Content = "Updated Content"
	updatedDiary, err := UpdateUserDiary(db, diary)
	assert.Nil(t, err)
	assert.Equal(t, updatedDiary.Question, "Updated Question")
	assert.Equal(t, updatedDiary.Content, "Updated Content")
}

// Test GetUserDiary
func TestGetUserDiary(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Insert a test diary
	diary := &UserDiary{UserID: 1, Question: "Test Question", Content: "Test Content"}
	err = CreateDiary(db, diary)
	assert.Nil(t, err)

	// Test GetUserDiary
	fetchedDiary, err := GetUserDiary(db, diary.DiaryID)
	assert.Nil(t, err)
	assert.Equal(t, fetchedDiary.DiaryID, diary.DiaryID)
	assert.Equal(t, fetchedDiary.Question, "Test Question")
}

// Test DeleteDiary
func TestDeleteDiary(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Insert a test diary
	diary := &UserDiary{UserID: 1, Question: "Delete Test Question", Content: "Delete Test Content"}
	err = CreateDiary(db, diary)
	assert.Nil(t, err)

	// Test DeleteDiary
	err = DeleteDiary(db, diary.DiaryID)
	assert.Nil(t, err)

	// Ensure the diary was deleted
	_, err = GetUserDiary(db, diary.DiaryID)
	assert.NotNil(t, err) // Expect error since diary should be deleted
}
