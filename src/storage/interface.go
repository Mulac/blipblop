package storage

import (
	"gorm.io/gorm"
)

// Job is, for now, an indeedJob
type Job IndeedJob

type IndeedJob struct {
	gorm.Model
	Title        string `gorm:"not null" json:"positionName"`
	Metadata     string `gorm:"not null" json:"metadata"`
	Description  string `gorm:"not null" json:"description"`
	TimePosted   string `gorm:"not null" json:"postedAt"`
	IndeedID     string `gorm:"primaryKey" json:"id"`
	IndeedURL    string `json:"url"`
	Location     string `gorm:"not null" json:"location"`
	Company      string `json:"company"`
	CompanyImage string `json:"companyImage"`
}

type JobDatabase interface {
	GetJobs() ([]Job, error)
	AddJob(...Job) error
}
