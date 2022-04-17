package cache

import (
	"database/sql"
	"log"

	"github.com/jpawlowskii/TechnikInformatykBackend/db"
)

func (c *Cache) UpdateCache(backendDatabase *sql.DB) {
	c.Database = backendDatabase
	if isCacheEnabled() {
		c.Exams = db.GetExams(backendDatabase)
		c.Questions = db.GetQuestions(backendDatabase)
		c.Images = map[string][]byte{}
		for i := range c.Questions {
			if c.Questions[i].HaveImage {
				c.Images[c.Questions[i].Uuid] = db.GetImage(backendDatabase, c.Questions[i].Uuid)
			}
		}
		c.DatabaseVersion = db.CalculateDatabaseVersion()
		log.Println("Cache updated")
	} else {
		log.Println("Cache is disabled in config. Ignoring update")
	}

}
