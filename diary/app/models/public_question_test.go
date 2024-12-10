package models

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPublicQuestion(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Insert test data
	questions := []PublicQuestion{
		{UserID: 1, Content: "Question 1", Likes: 5, Dislikes: 1},
		{UserID: 2, Content: "Question 2", Likes: 8, Dislikes: 2},
		{UserID: 3, Content: "Question 3", Likes: 10, Dislikes: 3},
	}
	for _, q := range questions {
		err := CreatePublicQuestion(db, &q)
		if err != nil {
			t.Fatalf("Failed to create public question: %v", err)
		}
	}

	// Test pagination: retrieve questions after ID 1
	lastID := 1
	pageSize := 2
	result, err := GetPublicQuestion(db, lastID, pageSize)
	if err != nil {
		t.Fatalf("Failed to get public questions: %v", err)
	}
	// Assert that we get the correct number of questions
	assert.Len(t, result, pageSize, "Expected to get 2 questions")

	// Assert that the questions have IDs greater than the last ID
	assert.Greater(t, result[0].ID, lastID, "Expected questions to have IDs greater than the last ID")
}

// TestCreatePublicQuestion tests the CreatePublicQuestion function
func TestCreatePublicQuestion(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Create a new public question
	q := &PublicQuestion{
		UserID:  1,
		Content: "What is the meaning of life?",
		Likes:   10,
	}

	err = CreatePublicQuestion(db, q)
	if err != nil {
		t.Fatalf("Failed to create public question: %v", err)
	}

	// Assert that the question is created with a non-zero ID
	assert.NotZero(t, q.ID, "ID should be auto-generated")
}

// TestInfiniteScroll tests the GetPublicQuestion function with multiple pages of data
func TestInfiniteScroll(t *testing.T) {
	db, err := SetupTestDB()
	assert.Nil(t, err)

	// Insert 5 test questions
	for i := 1; i <= 5; i++ {
		q := &PublicQuestion{
			UserID:  i,
			Content: "Question " + strconv.Itoa(i),
			Likes:   i * 2,
		}
		err := CreatePublicQuestion(db, q)
		if err != nil {
			t.Fatalf("Failed to create public question: %v", err)
		}
	}

	// Test pagination: first page, 2 questions per page
	lastID := 0
	pageSize := 2
	for page := 1; page <= 3; page++ {
		result, err := GetPublicQuestion(db, lastID, pageSize)
		if err != nil {
			t.Fatalf("Failed to get public questions: %v", err)
		}

		// Assert that the result contains the correct number of questions
		assert.True(t, len(result) <= pageSize, "Expected up to 2 questions per page, but got %d", len(result))

		// Update lastID for the next page
		lastID = result[len(result)-1].ID
	}
}
