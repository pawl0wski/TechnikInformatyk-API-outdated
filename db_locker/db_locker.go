/*
DB locker is an object that allows you to supervise the order of database queries.
This avoids errors when multiple requests are sent at once.
*/

package dblocker

import (
	"database/sql"
	"sync"
)

var lock = &sync.Mutex{}

type DBLocker struct {
	DB *sql.DB
}

var singleInstance *DBLocker

func GetDBLockerInstance() *DBLocker {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &DBLocker{}
		}
	}
	return singleInstance
}
