package storage

import "gorm.io/gorm"

// Job is, for now, an indeedJob
type Job IndeedJob

type IndeedJob struct {
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
	GetJobs() ([]Job, error)
	AddJob(Job) error
}
