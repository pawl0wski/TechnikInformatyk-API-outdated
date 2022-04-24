package cache

import (
	"database/sql"
	"sync"

	"github.com/pawl0wski/technikinformatyk-backend/model"
)

var lock = &sync.Mutex{}

type Cache struct {
	Exams           []model.Exam
	Questions       []model.Question
	Images          map[string][]byte
	CacheEnabled    bool
	Database        *sql.DB
	DatabaseVersion model.DatabaseVersion
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
