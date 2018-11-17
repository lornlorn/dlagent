package models

import (
	"app/db"
	"errors"
	"log"
)

/*
SysWorkflowInf struct map to table sys_workflow_inf
*/
type SysWorkflowInf struct {
	WfiId      int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	WfiName    string `xorm:"VARCHAR(128) NOT NULL"`
	WfiDesc    string `xorm:"VARCHAR(1024)"`
	WfiStatus  string `xorm:"VARCHAR(8) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}

/*
GetWorkflows func() ([]SysWorkflowInf, error)
*/
func GetWorkflows() ([]SysWorkflowInf, error) {

	workflows := make([]SysWorkflowInf, 0)

	// if err := db.Engine.Where("job_type = ?", jobType).Find(&jobs); err != nil {
	if err := db.Engine.Find(&workflows); err != nil {
		// return nil, err
		log.Printf("models.workflow_inf.GetWorkflows -> db.Engine.Find Error : %v\n", err)
		return nil, err
	}

	// for i, v := range crons {
	//  log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return workflows, nil
}

/*
GetWorkflowByID func(wfid int) (SysWorkflowInf, error)
*/
func GetWorkflowByID(wfid int) (SysWorkflowInf, error) {

	wf := new(SysWorkflowInf)
	wf.WfiId = wfid

	has, err := db.Engine.Get(wf)
	if err != nil {
		log.Printf("models.workflow_inf.GetWorkflowByID -> db.Engine.Get Error : %v\n", err)
		return SysWorkflowInf{}, err
	}

	if !has {
		return SysWorkflowInf{}, errors.New("Get 0 rows")
	}

	log.Println(wf)

	// for i, v := range crons {
	//  log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return *wf, nil
}
