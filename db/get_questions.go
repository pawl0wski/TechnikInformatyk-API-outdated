package db

import (
	"database/sql"
	"log"

	"github.com/jpawlowskii/TechnikInformatykBackend/structs"
)

func GetExamsForQuestion(backendDatabase *sql.DB, questionUuid string) []string {
	rows, err := backendDatabase.Query("SELECT examUUID FROM questionToExam WHERE questionUUID = ?;", questionUuid)
	if err != nil {
		log.Fatalln(err)
	}
	examUuids := []string{}
	for rows.Next() {
		var examUuid string
		rows.Scan(&examUuid)
		examUuids = append(examUuids, examUuid)
	}
	return examUuids
}

func GetQuestions(backendDatabase *sql.DB) []structs.Question {
	questions := []structs.Question{}
	rows, err := backendDatabase.Query("SELECT uuid, content, IF(image is not NULL, 1, 0) as haveImage, answerA, answerB, answerC, answerD, correct FROM question")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			uuid      string
			content   string
			haveImage int8
			answerA   string
			answerB   string
			answerC   string
			answerD   string
			correct   int8
		)

		question := structs.Question{}

		err := rows.Scan(&uuid, &content, &haveImage, &answerA, &answerB, &answerC, &answerD, &correct)
		if err != nil {
			log.Fatalln(err)
		}

		question.Uuid = uuid
		question.Content = content
		question.HaveImage = (haveImage == 1)
		question.ExamUuids = GetExamsForQuestion(backendDatabase, uuid)

		questions = append(questions, question)
	}
	return questions
}
