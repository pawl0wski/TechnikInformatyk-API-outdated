package dblocker

import (
	"sync"

	"github.com/pawl0wski/technikinformatyk-backend/db"
	"github.com/pawl0wski/technikinformatyk-backend/model"
)

var getterLock = &sync.Mutex{}

func (c DBLocker) GetExams() []model.Exam {
	getterLock.Lock()
	defer getterLock.Unlock()
	return db.GetExams(c.DB)
}

func (c DBLocker) GetQuestions() ([]model.Question, error) {
	getterLock.Lock()
	defer getterLock.Unlock()
	questions, err := db.GetQuestions(c.DB)
	if err != nil {
		return nil, err
	}
	return questions, err
}

func (c DBLocker) GetImage(questionUuid string) []byte {
	getterLock.Lock()
	defer getterLock.Unlock()
	return db.GetImage(c.DB, questionUuid)
}

func (c DBLocker) GetDatabaseVersion() model.DatabaseVersion {
	getterLock.Lock()
	defer getterLock.Unlock()
	return db.CalculateDatabaseVersion(c.DB)
}
