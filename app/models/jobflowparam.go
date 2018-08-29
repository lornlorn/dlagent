package models

import (
	"app/db"
	"errors"
	"log"
)

/*
JobflowParam struct map table jobflow_param
*/
type JobflowParam struct {
	JfpId         int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	JfpJfId       int    `xorm:"INTEGER NOT NULL"`
	JfpSeq        int    `xorm:"INTEGER NOT NULL"`
	JfpParameter  string `xorm:"VARCHAR(64) NOT NULL"`
	JfpDefault    string `xorm:"VARCHAR(512)"`
	JfpStatus     string `xorm:"VARCHAR(16) NOT NULL"`
	JfpRemark     string `xorm:"VARCHAR(512)"`
	JfpCreate     string `xorm:"VARCHAR(32)"`
	JfpCreatetime string `xorm:"VARCHAR(15)"`
	JfpModify     string `xorm:"VARCHAR(32)"`
	JfpModifytime string `xorm:"VARCHAR(15)"`
}

/*
GetJobFlowParamsByJobFlowID func(jobFlowIDs int) ([]JobflowParam, error)
*/
func GetJobFlowParamsByJobFlowID(jobFlowID int) ([]JobflowParam, error) {

	jobFlowParams := make([]JobflowParam, 0)

	if err := db.Engine.Where("jfp_jf_id = ?", jobFlowID).Find(&jobFlowParams); err != nil {
		log.Println(err)
		return []JobflowParam{}, err
	}

	if len(jobFlowParams) == 0 {
		return nil, errors.New("models.jobflowparam.GetJobFlowParamsByJobFlowID : No Records")
	}

	return jobFlowParams, nil
}
