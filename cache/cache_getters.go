package cache

import (
	"sync"

	"github.com/pawl0wski/TechnikInformatykBackend/db"
	"github.com/pawl0wski/TechnikInformatykBackend/model"
)

var getterLock = &sync.Mutex{}

func (c Cache) GetExams() []model.Exam {
	getterLock.Lock()
	defer getterLock.Unlock()
	if isCacheEnabled() {
		return c.Exams
	} else {
		return db.GetExams(c.Database)
	}
}

func (c Cache) GetQuestions() ([]model.Question, error) {
	getterLock.Lock()
	defer getterLock.Unlock()
	if isCacheEnabled() {
		return c.Questions, nil
	} else {
		questions, err := db.GetQuestions(c.Database)
		if err != nil {
			return nil, err
		}
		return questions, err
	}
}

func (c Cache) GetImage(questionUuid string) []byte {
	getterLock.Lock()
	defer getterLock.Unlock()
	if isCacheEnabled() {
		return c.Images[questionUuid]
	} else {
		return db.GetImage(c.Database, questionUuid)
	}
}

func (c Cache) GetDatabaseVersion() model.DatabaseVersion {
	getterLock.Lock()
	defer getterLock.Unlock()
	if isCacheEnabled() {
		return c.DatabaseVersion
	} else {
		return db.CalculateDatabaseVersion(c.Database)
	}
}
