package storage

import (
	"fmt"

	"gorm.io/gorm"
)

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

func (j *Job) FromIndeedResponse(resp map[string]interface{}) error {
	// TODO(jj) - add type check validation
	j.Title = fmt.Sprint(resp["positionName"])
	j.Company = fmt.Sprint(resp["company"])
	j.Location = fmt.Sprint(resp["location"])
	j.IndeedURL = fmt.Sprint(resp["url"])
	j.IndeedID = fmt.Sprint(resp["id"])
	j.Description = fmt.Sprint(resp["description"])
	j.Metadata = fmt.Sprint(resp["metadata"])
	return nil
}

type JobDatabase interface {
	GetJobs() ([]Job, error)
	AddJob(Job) error
}
