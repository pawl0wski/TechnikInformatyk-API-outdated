package db

import (
	"database/sql"
	"io"
	"testing"

	"github.com/cockroachdb/copyist"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pawl0wski/technikinformatyk-backend/model"
)

func init() {
	copyist.Register("mysql")
}

func getDb(t *testing.T) (*sql.DB, io.Closer) {
	connPath := GetDbConnectionPath()

	ioCloser := copyist.Open(t)

	db, _ := sql.Open("copyist_mysql", connPath)
	return db, ioCloser
}

func TestGetExams(t *testing.T) {
	db, closer := getDb(t)
	defer closer.Close()

	exams := GetExams(db)
	if len(exams) == 0 {
		t.Errorf("GetExams result is empty")
	}

	for i, v := range exams {
		if v == (model.Exam{}) {
			t.Errorf("Exam %d is empty", i)
		}
	}
}

func TestGetQuestions(t *testing.T) {
	db, closer := getDb(t)
	defer closer.Close()

	questions, err := GetQuestions(db)
	if err != nil {
		t.Errorf("GetQuestions returned error \"%s\"", err)
	}

	if len(questions) == 0 {
		t.Errorf("GetQuestions result is empty")
	}

	for i, v := range questions {
		if v.IsEmpty() {
			t.Errorf("Question %d is empty", i)
		}
	}
}

// FIXME: 2022/05/07 15:54:47 driver: skip fast-path; continue as if unimplemented
// func TestGetImage(t *testing.T) {
// 	db, closer := getDb(t)
// 	defer closer.Close()

// 	QuestionUUIDWithImage := "008ac2ea-f5f9-46d8-900d-a17cf0b94a80"
// 	QuestionUUIDWithoutImage := "000e457f-8e2e-4265-a8e2-ddd2bc92e5d3"

// 	image := GetImage(db, QuestionUUIDWithImage)
// 	if len(image) == 0 {
// 		t.Errorf("Question with image returned no image")
// 	}

// 	image = GetImage(db, QuestionUUIDWithoutImage)
// 	if !(len(image) == 0) {
// 		t.Errorf("Question without image returned image")
// 	}

// }

func TestGetDbVersion(t *testing.T) {
	db, closer := getDb(t)
	defer closer.Close()

	dbVersion := CalculateDatabaseVersion(db)
	if dbVersion.Version == 0 {
		t.Error("DB version is 0")
	}
}
