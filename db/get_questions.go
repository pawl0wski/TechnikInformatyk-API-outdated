package db

import (
	"database/sql"
	"log"
	"strings"

	"github.com/pawl0wski/TechnikInformatykBackend/structs"
)

func parseExamResponseFromDatabase(examUuids string) []string {
	return strings.Split(examUuids, ",")
}

func GetQuestions(backendDatabase *sql.DB) []structs.Question {
	questions := []structs.Question{}
	rows, err := backendDatabase.Query("SELECT uuid, content, IF(image is not NULL, 1, 0) as haveImage, answerA, answerB, answerC, answerD, correct, GROUP_CONCAT(questionToExam.examUUID) as examUuids FROM question INNER JOIN questionToExam on questionToExam.questionUUID = question.uuid GROUP BY question.uuid;")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var examUuids string
		question := structs.Question{}

		err := rows.Scan(&question.Uuid, &question.Content, &question.HaveImage, &question.AnswerA, &question.AnswerB, &question.AnswerC, &question.AnswerD, &question.CorrectAnswer, &examUuids)
		if err != nil {
			log.Fatalln(err)
		}

		question.ExamUuids = parseExamResponseFromDatabase(examUuids)

		questions = append(questions, question)
	}
	return questions
}
