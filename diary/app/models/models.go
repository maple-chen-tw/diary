package models

import (
	"time"
)

// User 代表 users 表中的用戶。
type User struct {
	UserID       int       `json:"user_id" db:"user_id" gorm:"primary_key;auto_increment"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	Email        string    `json:"email" db:"email"`
	CreateAt     time.Time `json:"create_at" db:"create_at"`
}

// DailyQuestion 代表 daily_questions 表中的每日問題。
type DailyQuestion struct {
	QuestionID int    `json:"question_id" db:"question_id" gorm:"primaryKey;autoIncrement"`
	Question   string `json:"question" db:"question"`
}

// UserDiary 代表 user_diaries 表中的日記條目。
// question_id 和 question 都存在，避免每次query都要做Join。
type UserDiary struct {
	DiaryID    int       `json:"diary_id" db:"diary_id" gorm:"primary_key;auto_increment"`
	UserID     int       `json:"user_id" db:"user_id"`
	QuestionID int       `json:"question_id" db:"question_id"`
	Question   string    `json:"question" db:"question"`
	Content    string    `json:"content" db:"content"`
	CreateAt   time.Time `json:"create_at" db:"create_at"`
	UpdateAt   time.Time `json:"update_at" db:"update_at"`
}

//---

// PublicQuestion 代表 public_questions 表中的公開問題。
type PublicQuestion struct {
	ID     int `json:"id" db:"id" gorm:"primary_key;auto_increment"`
	UserID int `json:"user_id" db:"user_id"`
	//Title    string    `json:"title" db:"title"`
	Content  string    `json:"content" db:"content"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
	Likes    int       `json:"likes" db:"likes"`
	Dislikes int       `json:"dislikes" db:"dislikes"`
	//UpdateAt time.Time `json:"update_at" db:"update_at"`
}

// PublicQuestionAnswer 代表 public_questions_answer 表中的回復。
type PublicQuestionAnswer struct {
	ID         int       `json:"id" db:"id" gorm:"primary_key;auto_increment"`
	QuestionID int       `json:"question_id" db:"question_id"`
	UserID     int       `json:"user_id" db:"user_id"`
	Content    string    `json:"content" db:"content"`
	CreateAt   time.Time `json:"create_at" db:"create_at"`
}

// Vote 代表 votes 表中的投票。
type Vote struct {
	ID          int    `json:"id" db:"id" gorm:"primary_key;auto_increment"`
	ContentID   int    `json:"content_id" db:"content_id"`
	UserID      int    `json:"user_id" db:"user_id"`
	ContentType string `json:"content_type" db:"content_type"`
	VoteType    string `json:"vote_type" db:"vote_type"`
}
