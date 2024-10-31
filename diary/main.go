package main

import (
	model "diary/model"
	route "diary/route"
	"log"
)

var testUsers []model.User
var testDailyQuestions []model.DailyQuestion
var testUserDiaries []model.UserDiary
var testPublicQuestions []model.PublicQuestion
var testPublicAnswers []model.PublicQuestionAnswer
var testVotes []model.Vote

func InitializeTestData() {
	testUsers, testDailyQuestions, testUserDiaries, testPublicQuestions, testPublicAnswers, testVotes = model.GenerateTestData()
}

func main() {
	InitializeTestData()

	router := route.SetupRouter()
	log.Println("Server is running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
