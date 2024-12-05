package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test CreateDailyQuestion
func TestCreateDailyQuestion(t *testing.T) {

	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Create a new DailyQuestion instance
	dailyQuestion := &DailyQuestion{
		Question: "What is the capital of France?",
	}

	// Call the CreateDailyQuestion function
	err = CreateDailyQuestion(db, dailyQuestion)

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert the record was created with the correct values
	var createdDailyQuestion DailyQuestion
	db.First(&createdDailyQuestion, "question = ?", "What is the capital of France?")
	assert.Equal(t, "What is the capital of France?", createdDailyQuestion.Question)
	assert.NotZero(t, createdDailyQuestion.QuestionID, "QuestionID should be auto-generated")
}

// Test GetDailyQuestionByID
func TestGetDailyQuestionByID(t *testing.T) {
	// Set up the test database
	db, err := SetupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test DB: %v", err)
	}

	// Create a new DailyQuestion instance
	dailyQuestion := &DailyQuestion{
		Question: "What is the largest ocean on Earth?",
	}

	// Save the question to the database
	err = CreateDailyQuestion(db, dailyQuestion)
	if err != nil {
		t.Fatalf("failed to create daily question: %v", err)
	}

	// Retrieve the question by ID
	retrievedQuestion, err := GetDailyQuestionByID(db, dailyQuestion.QuestionID)

	// Assert no error occurred and the correct question is retrieved
	assert.NoError(t, err)
	assert.Equal(t, "What is the largest ocean on Earth?", retrievedQuestion.Question)
}

// Test GetDailyQuestionByRandom
func TestGetDailyQuestionByRandom(t *testing.T) {
	// Set up the test database
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Insert multiple questions
	questions := []DailyQuestion{
		{Question: "What is the largest country by area?"},
		{Question: "What is the smallest planet in the solar system?"},
		{Question: "What is the longest river in the world?"},
	}
	for _, q := range questions {
		err := CreateDailyQuestion(db, &q)
		if err != nil {
			t.Fatalf("failed to create daily question: %v", err)
		}
	}

	// Retrieve a random question
	randomQuestion, err := GetDailyQuestionByRandom(db)

	// Assert no error occurred and a valid question is retrieved
	assert.NoError(t, err)
	assert.Contains(t, []string{
		"What is the largest country by area?",
		"What is the smallest planet in the solar system?",
		"What is the longest river in the world?",
	}, randomQuestion.Question)
}
