package storage

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type jobDatabaseImpl struct {
	*gorm.DB
}

func (db jobDatabaseImpl) GetJobs() ([]Job, error) {
	jobs := []Job{}
	result := db.Find(&jobs)
	if result.Error != nil {
		return []Job{}, fmt.Errorf("ERROR|jobDatabaseImpl.GetJobs()|%v", result.Error)
	}

	return jobs, nil
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
	dsn := "root@tcp(127.0.0.1:3306)/prototype?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// This updates our tables to match the schema defined by the struct
	db.AutoMigrate(&Job{})

	return &jobDatabaseImpl{db}
}
