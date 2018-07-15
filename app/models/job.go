package models

import (
	"app/db"
	"log"
)

/*
Job struct map table job
*/
type Job struct {
	JobId         int    `xorm:"INTEGER not null pk"` // xorm
	JobName       string `xorm:"VARCHAR not null"`
	JobType       string `xorm:"VARCHAR not null"`
	JobCrontime   string `xorm:"VARCHAR"`
	JobPlantime   string `xorm:"VARCHAR"`
	JobStatus     string `xorm:"VARCHAR not null"`
	JobRemark     string `xorm:"VARCHAR"`
	JobCreate     string `xorm:"VARCHAR"`
	JobCreatetime string `xorm:"VARCHAR"`
	JobModify     string `xorm:"VARCHAR"`
	JobModifytime string `xorm:"VARCHAR"`
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
GetJobList func() ([]Job, error)
*/
func GetJobList() ([]Job, error) {

	jobs := make([]Job, 0)
	// if err := db.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	if err := db.Engine.Where("job_status = ?", "READY").Find(&jobs); err != nil {
		// return nil, err
		log.Println(err)
		return nil, err
	}

	// for i, v := range crons {
	// 	log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return jobs, nil
}
