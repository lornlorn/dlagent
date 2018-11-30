package models

import (
	"app/utils"
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

	// if err := utils.Engine.Where("job_type = ?", jobType).Find(&jobs); err != nil {
	if err := utils.Engine.Find(&workflows); err != nil {
		// return nil, err
		log.Printf("models.workflow_inf.GetWorkflows -> utils.Engine.Find Error : %v\n", err)
		return nil, err
	}

	// for i, v := range crons {
	//  log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return workflows, nil
}

/*
GetWorkflowByID func(wfiid int) (SysWorkflowInf, error)
*/
func GetWorkflowByID(wfiid int) (SysWorkflowInf, error) {

	wf := new(SysWorkflowInf)
	wf.WfiId = wfiid

	has, err := utils.Engine.Get(wf)
	if err != nil {
		log.Printf("models.workflow_inf.GetWorkflowByID -> utils.Engine.Get Error : %v\n", err)
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

/*
DelWorkflowByID func(wfiid int)
*/
func DelWorkflowByID(wfiid int) error {
	wfi := new(SysWorkflowInf)
	wfi.WfiId = wfiid
	affected, err := utils.Engine.Delete(wfi)
	if err != nil {
		log.Printf("models.workflow_inf.DelWorkflowByID -> utils.Engine.Delete Error : %v\n", err)
		return err
	}
	log.Printf("删除%v条记录\n", affected)
	return nil
}
