package db

import (
	"database/sql"
	"log"

	"github.com/pawl0wski/TechnikInformatykBackend/model"
)

func GetExams(backendDatabase *sql.DB) []model.Exam {
	row, err := backendDatabase.Query("SELECT uuid, name, description, icon FROM exam")
	if err != nil {
		log.Fatalln(err)
	}
	defer row.Close()
	var exams []model.Exam = []model.Exam{}
	for row.Next() {
		var exam model.Exam
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
