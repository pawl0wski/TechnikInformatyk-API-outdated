package cache

import (
	"database/sql"
	"log"
	"sync"

	"github.com/Jeboczek/TechnikInformatykBackend/db"
	"github.com/Jeboczek/TechnikInformatykBackend/structs"
)

var lock = &sync.Mutex{}

type Cache struct {
	Exams     []structs.Exam
	Questions []structs.Question
}

var singleInstance *Cache

func GetCacheInstance() *Cache {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Cache{}
		}
	}
	return singleInstance
}

func (c Cache) GetExams() []structs.Exam {
	return c.Exams
}

func (c Cache) GetQuestions() []structs.Question {
	return c.Questions
}

func (c *Cache) UpdateCache(backendDatabase *sql.DB) {
	c.Exams = db.GetExams(backendDatabase)
	c.Questions = db.GetQuestions(backendDatabase)
	log.Println("Cache updated")
}
