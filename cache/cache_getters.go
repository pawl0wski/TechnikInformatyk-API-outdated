package cache

import (
	"github.com/jpawlowskii/TechnikInformatykBackend/db"
	"github.com/jpawlowskii/TechnikInformatykBackend/structs"
)

func (c Cache) GetExams() []structs.Exam {
	if isCacheEnabled() {
		return c.Exams
	} else {
		return db.GetExams(c.Database)
	}
}

func (c Cache) GetQuestions() []structs.Question {
	if isCacheEnabled() {
		return c.Questions
	} else {
		return db.GetQuestions(c.Database)
	}
}

func (c Cache) GetImage(questionUuid string) []byte {
	if isCacheEnabled() {
		return c.Images[questionUuid]
	} else {
		return db.GetImage(c.Database, questionUuid)
	}
}

func (c Cache) GetDatabaseVersion() structs.DatabaseVersion {
	if isCacheEnabled() {
		return c.DatabaseVersion
	} else {
		return db.CalculateDatabaseVersion()
	}
}
