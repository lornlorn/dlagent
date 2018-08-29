package models

import (
	"app/db"
	"errors"
	"log"
)

/*
Job struct map table job
*/
type Job struct {
	JobId       int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	JobName     string `xorm:"VARCHAR(64) NOT NULL"`
	JobType     string `xorm:"VARCHAR(16) NOT NULL"`
	JobCrontime string `xorm:"VARCHAR(128)"`
	JobPlantime string `xorm:"VARCHAR(15)"`
	JobStatus   string `xorm:"VARCHAR(16) NOT NULL"`
	JobRemark   string `xorm:"VARCHAR(512)"`
	// JobCreate     string `xorm:"VARCHAR(32)"`
	// JobCreatetime string `xorm:"VARCHAR(15)"`
	// JobModify     string `xorm:"VARCHAR(32)"`
	// JobModifytime string `xorm:"VARCHAR(15)"`
}

// Save insert method
func (j *Job) Save() error {
	// affected, err := db.Engine.Insert(d)
	_, err := db.Engine.Insert(j)
	if err != nil {
		return err
	}
	return nil
}

/*
GetJobs func(jobType string) ([]Job, error)
*/
func GetJobs(jobType string) ([]Job, error) {

	jobs := make([]Job, 0)

	if err := db.Engine.Where("job_type = ?", jobType).Find(&jobs); err != nil {
		// return nil, err
		log.Println(err)
		return nil, err
	}

	// for i, v := range crons {
	//  log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return jobs, nil
}

/*
GetJobByID func(jobid int) (Job, error)
*/
func GetJobByID(jobID int) (Job, error) {

	job := new(Job)
	job.JobId = jobID

	has, err := db.Engine.Get(job)
	if err != nil {
		log.Println(err)
		return Job{}, err
	}

	if !has {
		return Job{}, errors.New("Get 0 rows")
	}

	log.Println(job)

	// for i, v := range crons {
	//  log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return *job, nil
}
