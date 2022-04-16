package db

import (
	"database/sql"
	"log"

	"github.com/Jeboczek/TechnikInformatykBackend/structs"
)

func GetExams(backendDatabase *sql.DB, alsoReadQuestion bool) []structs.Exam {
	row, err := backendDatabase.Query("SELECT uuid, name, description, icon FROM exam")
	if err != nil {
		log.Fatalln(err)
	}
	defer row.Close()
	var exams []structs.Exam = []structs.Exam{}
	for row.Next() {
		var (
			uuid        string
			name        string
			description string
			icon        string
		)
		// Scan the row into the variables declared above.
		err := row.Scan(&uuid, &name, &description, &icon)

		// If there is an error, print it out and return.
		if err != nil {
			log.Fatalln(err)
		}

		// Get questions
		questions := []structs.Question{}

		// If alsoReadQuestion is true, get questions
		if alsoReadQuestion {
			questions = GetQuestions(backendDatabase, uuid)
		}

		// Create a new exam and append it to the slice.
		exam := structs.Exam{Uuid: uuid, Name: name, Description: description, Icon: icon, Questions: questions}
		exams = append(exams, exam)
	}
	return exams
}
