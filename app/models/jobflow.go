package models

import (
	"app/db"
	"errors"
	"log"
)

/*
Jobflow struct map table jobflow
*/
type Jobflow struct {
	JfId         int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	JfJobId      int    `xorm:"INTEGER NOT NULL"`
	JfName       string `xorm:"VARCHAR(128) NOT NULL"`
	JfSeq        int    `xorm:"INTEGER NOT NULL"`
	JfSh         string `xorm:"VARCHAR(64)"`
	JfCmd        string `xorm:"VARCHAR(256) NOT NULL"`
	JfStatus     string `xorm:"VARCHAR(16) NOT NULL"`
	JfRemark     string `xorm:"VARCHAR(512)"`
	JfCreate     string `xorm:"VARCHAR(32)"`
	JfCreatetime string `xorm:"VARCHAR(15)"`
	JfModify     string `xorm:"VARCHAR(32)"`
	JfModifytime string `xorm:"VARCHAR(15)"`
}

/*
GetJobFlowsByJobID func(jobid int) ([]Jobflow, error)
*/
func GetJobFlowsByJobID(jobid int) ([]Jobflow, error) {

	jobFlows := make([]Jobflow, 0)

	if err := db.Engine.Where("jf_job_id = ?", jobid).Find(&jobFlows); err != nil {
		log.Println(err)
		return []Jobflow{}, err
	}

	if len(jobFlows) == 0 {
		return nil, errors.New("models.jobflow.GetJobFlowsByJobID : No Records")
	}

	return jobFlows, nil
}
