package db

import "gorm.io/gorm"

// Job is, for now, an indeedJob
type Job indeedJob

type indeedJob struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Metadata     string `gorm:"not null"`
	Description  string `gorm:"not null"`
	TimePosted   string `gorm:"not null"`
	IndeedID     string
	IndeedURL    string
	Location     string `gorm:"not null"`
	Company      string
	CompanyImage string
}

type JobDatabase interface {
	GetJob(uint) (Job, error)
	AddJob(Job) error
}

func NewJobDatabase() JobDatabase {
	return newJobDatabase()
}
