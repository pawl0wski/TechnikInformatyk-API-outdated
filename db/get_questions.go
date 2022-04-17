package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	"github.com/jpawlowskii/TechnikInformatykBackend/structs"
)

type databaseQuestionAnswers struct {
	Content string `json:"content"`
	Correct bool   `json:"correct"`
}

func convertDatabaseAnswersToQuestionAnswers(answers string, question *structs.Question) {
	databaseAnswers := [4]databaseQuestionAnswers{}
	json.Unmarshal([]byte(answers), &databaseAnswers)

	// Convert database answers to question answers
	question.AnswerA = databaseAnswers[0].Content
	question.AnswerB = databaseAnswers[1].Content
	question.AnswerC = databaseAnswers[2].Content
	question.AnswerD = databaseAnswers[3].Content
	// Convert database correct answer to question correct answer
	for i := 0; i < 4; i++ {
		if databaseAnswers[i].Correct {
			question.CorretAnswer = int8(i)
		}
	}
}

func getUuidsFromCell(cell string) []string {
	uuids := []string{}
	if strings.Contains(cell, ",") {
		uuids = strings.Split(cell, ",")
	} else {
		uuids = append(uuids, cell)
	}
	return uuids
}

func GetQuestions(backendDatabase *sql.DB) []structs.Question {
	questions := []structs.Question{}
	rows, err := backendDatabase.Query("SELECT uuid, content, IIF(img is not NULL, 1, 0) as haveImage, answers, examUuids FROM question")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			uuid           string
			content        string
			haveImage      int8
			databaseAnswer string
			examUuids      string
		)

		question := structs.Question{}

		err := rows.Scan(&uuid, &content, &haveImage, &databaseAnswer, &examUuids)
		if err != nil {
			log.Fatalln(err)
		}

		question.Uuid = uuid
		question.Content = content
		question.HaveImage = (haveImage == 1)
		convertDatabaseAnswersToQuestionAnswers(databaseAnswer, &question)
		question.ExamUuids = getUuidsFromCell(examUuids)

		questions = append(questions, question)
	}
	return questions
}
