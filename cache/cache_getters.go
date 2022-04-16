package cache

import (
	"github.com/Jeboczek/TechnikInformatykBackend/db"
	"github.com/Jeboczek/TechnikInformatykBackend/structs"
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
