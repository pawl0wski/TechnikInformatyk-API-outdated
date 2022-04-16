package cache

import (
	"database/sql"
	"sync"

	"github.com/Jeboczek/TechnikInformatykBackend/structs"
)

var lock = &sync.Mutex{}

type Cache struct {
	Exams           []structs.Exam
	Questions       []structs.Question
	Images          map[string][]byte
	CacheEnabled    bool
	Database        *sql.DB
	DatabaseVersion structs.DatabaseVersion
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
