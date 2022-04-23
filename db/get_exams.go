package db

import (
	"database/sql"
	"log"

	"github.com/jpawlowskii/TechnikInformatykBackend/structs"
)

func GetExams(backendDatabase *sql.DB) []structs.Exam {
	row, err := backendDatabase.Query("SELECT uuid, name, description, icon FROM exam")
	if err != nil {
		log.Fatalln(err)
	}
	defer row.Close()
	var exams []structs.Exam = []structs.Exam{}
	for row.Next() {
		var exam structs.Exam
		// Scan the row into the exam declared above.
		err := row.Scan(&exam.Uuid, &exam.Name, &exam.Description, &exam.Icon)

		// If there is an error, print it out and return.
		if err != nil {
			log.Fatalln(err)
		}

		// Create a new exam and append it to the slice.
		exams = append(exams, exam)
	}
	return exams
}
