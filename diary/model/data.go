package models

import (
	"time"
)

// GenerateTestData 生成一些測試資料
func GenerateTestData() ([]User, []DailyQuestion, []UserDiary, []PublicQuestion, []PublicQuestionAnswer, []Vote) {
	users := []User{
		{ID: 1, Username: "user1", PasswordHash: "hashed_password1", Email: "user1@example.com", CreateAt: time.Now()},
		{ID: 2, Username: "user2", PasswordHash: "hashed_password2", Email: "user2@example.com", CreateAt: time.Now()},
	}

	questions := []DailyQuestion{
		{ID: 1, Question: "What are you grateful for today?"},
		{ID: 2, Question: "What did you learn today?"},
	}

	diaries := []UserDiary{
		{ID: 1, QuestionID: 1, Content: "I am grateful for my family.", CreateAt: time.Now(), UpdateAt: time.Now()},
		{ID: 2, QuestionID: 2, Content: "I learned about Go routines.", CreateAt: time.Now(), UpdateAt: time.Now()},
	}

	publicQuestions := []PublicQuestion{
		{ID: 1, UserID: 1, Title: "How to learn Go?", Content: "What resources do you recommend?", CreateAt: time.Now(), Likes: 5, Dislikes: 0, UpdateAt: time.Now()},
	}

	publicAnswers := []PublicQuestionAnswer{
		{ID: 1, QuestionID: 1, UserID: 2, Content: "I recommend the Go by Example site.", CreateAt: time.Now()},
	}

	votes := []Vote{
		{ID: 1, ContentID: 1, UserID: 2, ContentType: "question", VoteType: "up"},
	}

	return users, questions, diaries, publicQuestions, publicAnswers, votes
}
