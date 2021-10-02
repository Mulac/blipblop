package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title        string
	Metadata     string
	Description  string
	TimePosted   string
	IndeedID     string
	IndeedURL    string
	Location     string
	Company      string
	CompanyImage string
}

type JobDatabase interface {
	GetJob(uint) (Job, error)
	AddJob(Job) error
}

type jobDatabaseImpl struct {
	*gorm.DB
}

func (db jobDatabaseImpl) GetJob(id uint) (Job, error) {
	job := &Job{}
	result := db.First(job, "id = ?", id)
	if result.Error != nil {
		panic(result.Error)
	}

	return *job, nil
}

func (db jobDatabaseImpl) AddJob(job Job) error {
	return nil
}

func NewJobDatabase() JobDatabase {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(127.0.0.1:3306)/prototype?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Job{})

	return jobDatabaseImpl{db}
}
