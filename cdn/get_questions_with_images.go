package cdn

import (
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
	"github.com/pawl0wski/technikinformatyk-backend/model"
)

func getQuestionsWithImages() []model.Question {
	dbLocker := dblocker.GetDBLockerInstance()
	questions, _ := dbLocker.GetQuestions()

	questionsToReturn := []model.Question{}
	for _, question := range questions {
		if question.HaveImage {
			questionsToReturn = append(questionsToReturn, question)
		}
	}
	return questionsToReturn
}
