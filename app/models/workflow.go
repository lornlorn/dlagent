package models

import (
	"app/utils"
	"errors"

	seelog "github.com/cihub/seelog"
)

/*
TbWorkflow struct map to table tb_workflow
*/
type TbWorkflow struct {
	WfId       int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	WfNo       string `xorm:"VARCHAR(20) NOT NULL UNIQUE"`
	WfName     string `xorm:"VARCHAR(128) NOT NULL"`
	WfDesc     string `xorm:"VARCHAR(1024)"`
	WfStatus   string `xorm:"VARCHAR(8) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}

/*
NewWorkflow struct map to table tb_workflow without column WfId
*/
type NewWorkflow struct {
	WfNo       string `xorm:"VARCHAR(20) NOT NULL UNIQUE"`
	WfName     string `xorm:"VARCHAR(128) NOT NULL"`
	WfDesc     string `xorm:"VARCHAR(1024)"`
	WfStatus   string `xorm:"VARCHAR(8) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}

/*
TableName xorm mapper
NewWorkflow struct map to table tb_workflow
*/
func (wf NewWorkflow) TableName() string {
	return "tb_workflow"
}

// Save insert method
func (wf NewWorkflow) Save() error {
	affected, err := utils.Engine.Insert(wf)
	if err != nil {
		// seelog.Errorf("utils.Engine.Insert Error : %v", err)
		return err
	}
	seelog.Debugf("%v insert : %v", affected, wf)

	return nil
}

// Update method
func (wf TbWorkflow) Update() error {
	affected, err := utils.Engine.ID(wf.WfId).Update(wf)
	if err != nil {
		// seelog.Errorf("utils.Engine.ID.Update Error : %v", err)
		return err
	}
	seelog.Debugf("%v update : %v", affected, wf)

	return nil
}

/*
GetWorkflows func() ([]TbWorkflow, error)
*/
func GetWorkflows() ([]TbWorkflow, error) {
	workflows := make([]TbWorkflow, 0)

	if err := utils.Engine.Find(&workflows); err != nil {
		// seelog.Errorf("utils.Engine.Find Error : %v", err)
		return nil, err
	}

	return workflows, nil
}

/*
GetWorkflowByID func(wfid int) (TbWorkflow, error)
*/
func GetWorkflowByID(wfid int) (TbWorkflow, error) {
	wf := new(TbWorkflow)
	wf.WfId = wfid

	has, err := utils.Engine.Get(wf)
	if err != nil {
		// seelog.Errorf("utils.Engine.Get Error : %v", err)
		return TbWorkflow{}, err
	}

	if !has {
		// seelog.Debug("Get 0 row")
		return TbWorkflow{}, errors.New("Get 0 row")
	}

	return *wf, nil
}

/*
DelWorkflowByID func(wfid int) error
*/
func DelWorkflowByID(wfid int) error {
	wf := new(TbWorkflow)
	wf.WfId = wfid

	affected, err := utils.Engine.Delete(wf)
	if err != nil {
		// seelog.Errorf("utils.Engine.Delete Error : %v", err)
		return err
	}
	seelog.Debugf("%v delete : %v", affected, wf)

	return nil
}
