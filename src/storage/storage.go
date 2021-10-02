package storage

import "sync"

var db JobDatabase
var once sync.Once

func DB() JobDatabase {
	once.Do(func() {
		db = newJobDatabase()
	})

	return db
}
