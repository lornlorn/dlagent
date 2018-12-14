package models

import (
	"app/utils"
	"errors"

	seelog "github.com/cihub/seelog"
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
NewWorkflowInf struct map to table sys_workflow_inf
*/
type NewWorkflowInf struct {
	WfiName    string `xorm:"VARCHAR(128) NOT NULL"`
	WfiDesc    string `xorm:"VARCHAR(1024)"`
	WfiStatus  string `xorm:"VARCHAR(8) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}

/*
TableName xorm mapper
NewWorkflowInf struct map to table sys_workflow_inf
*/
func (wfi NewWorkflowInf) TableName() string {
	return "sys_workflow_inf"
}

// Save insert method
func (wfi NewWorkflowInf) Save() error {
	affected, err := utils.Engine.Insert(wfi)
	if err != nil {
		seelog.Errorf("utils.Engine.Insert Error : %v", err)
		return err
	}
	seelog.Debugf("%v insert : %v", affected, wfi)
	return nil
}

// Update method
func (wfi SysWorkflowInf) Update() error {
	affected, err := utils.Engine.ID(wfi.WfiId).Update(wfi)
	if err != nil {
		seelog.Errorf("utils.Engine.ID.Update Error : %v", err)
		return err
	}
	seelog.Debugf("%v update : %v", affected, wfi)
	return nil
}

/*
GetWorkflows func() ([]SysWorkflowInf, error)
*/
func GetWorkflows() ([]SysWorkflowInf, error) {

	workflows := make([]SysWorkflowInf, 0)

	if err := utils.Engine.Find(&workflows); err != nil {
		seelog.Errorf("utils.Engine.Find Error : %v", err)
		return nil, err
	}
	seelog.Debugf("Workflows : %v", workflows)

	return workflows, nil
}

/*
GetWorkflowByID func(wfiid int) (SysWorkflowInf, error)
*/
func GetWorkflowByID(wfiid int) (SysWorkflowInf, error) {

	wfi := new(SysWorkflowInf)
	wfi.WfiId = wfiid

	has, err := utils.Engine.Get(wfi)
	if err != nil {
		seelog.Errorf("utils.Engine.Get Error : %v", err)
		return SysWorkflowInf{}, err
	}

	if !has {
		seelog.Debug("Get 0 row")
		return SysWorkflowInf{}, errors.New("Get 0 row")
	}

	seelog.Debugf("Workflow : %v", wfi)

	return *wfi, nil
}

/*
DelWorkflowByID func(wfiid int)
*/
func DelWorkflowByID(wfiid int) error {
	wfi := new(SysWorkflowInf)
	wfi.WfiId = wfiid

	affected, err := utils.Engine.Delete(wfi)
	if err != nil {
		seelog.Errorf("utils.Engine.Delete Error : %v", err)
		return err
	}
	seelog.Debugf("%v delete : %v", affected, wfi)

	return nil
}
