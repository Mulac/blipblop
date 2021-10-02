package db

import (
	. "blipblop/src/core/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type jobDatabaseImpl struct {
	*gorm.DB
}

func (db jobDatabaseImpl) GetJob(id uint) (Job, error) {
	job := &Job{}
	result := db.First(job, "id = ?", id)
	if result.Error != nil {
		return Job{}, fmt.Errorf("ERROR|jobDatabaseImpl.GetJob(%d)|%v", id, result.Error)
	}

	return *job, nil
}

func (db jobDatabaseImpl) AddJob(job Job) error {
	result := db.Create(&job)
	if result.Error != nil {
		return fmt.Errorf("ERROR|jobDatabaseImpl.AddJob(%+v)|%v", job, result.Error)
	}

	return nil
}

func newJobDatabase() *jobDatabaseImpl {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dbCfg := Config().Database
	conn, err := dbCfg.GetHost()
	if err != nil {
		log.Fatalf("ERROR|newJobDatabase|failed to get connection string|%v", err)
	}

	var dialect gorm.Dialector
	switch dbCfg.Driver {
	case "mysql":
		dialect = mysql.Open(conn)
	case "sqlite3":
		dialect = sqlite.Open(conn)
	}

	// dsn := "root@tcp(127.0.0.1:3306)/prototype?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// This updates our tables to match the schema defined by the struct
	db.AutoMigrate(&Job{})

	return &jobDatabaseImpl{db}
}
