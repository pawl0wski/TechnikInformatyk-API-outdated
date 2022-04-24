package db

import (
	"database/sql"
	"strings"

	"github.com/pawl0wski/TechnikInformatykBackend/model"
)

func parseExamResponseFromDatabase(examUuids string) []string {
	return strings.Split(examUuids, ",")
}

func GetQuestions(backendDatabase *sql.DB) ([]model.Question, error) {
	questions := []model.Question{}
	rows, err := backendDatabase.Query("SELECT uuid, content, IF(image is not NULL, 1, 0) as haveImage, answerA, answerB, answerC, answerD, correct, GROUP_CONCAT(questionToExam.examUUID) as examUuids FROM question INNER JOIN questionToExam on questionToExam.questionUUID = question.uuid GROUP BY question.uuid;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var examUuids string
		question := model.Question{}

		err := rows.Scan(&question.Uuid, &question.Content, &question.HaveImage, &question.AnswerA, &question.AnswerB, &question.AnswerC, &question.AnswerD, &question.CorrectAnswer, &examUuids)
		if err != nil {
			return nil, err
		}

		question.ExamUuids = parseExamResponseFromDatabase(examUuids)

		questions = append(questions, question)
	}
	return questions, nil
}
