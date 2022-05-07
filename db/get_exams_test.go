package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cockroachdb/copyist"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	copyist.Register("mysql")
}

func TestGetExams(t *testing.T) {
	connPath := GetDbConnectionPath()
	fmt.Println(connPath)

	defer copyist.Open(t).Close()

	db, _ := sql.Open("copyist_mysql", connPath)
	defer db.Close()

	exams := GetExams(db)
	examsExpected := 2
	if len(exams) != 2 {
		t.Errorf("Expected %d exams got %d", examsExpected, len(exams))
	}
}
