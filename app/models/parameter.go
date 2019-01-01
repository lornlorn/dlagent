package models

import (
	"app/utils"
	"errors"

	seelog "github.com/cihub/seelog"
)

/*
SysWorkflowParam struct map to table sys_workflow_param
*/
type SysWorkflowParam struct {
	WfpId      int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	WfdId      int    `xorm:"INTEGER NOT NULL"`
	WfpSeq     int    `xorm:"INTEGER NOT NULL"`
	WfpName    string `xorm:"VARCHAR(128)"`
	WfpDefault string `xorm:"VARCHAR(128) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}

/*
NewWorkflowParam struct map to table sys_workflow_param
*/
type NewWorkflowParam struct {
	WfdId      int    `xorm:"INTEGER NOT NULL"`
	WfpSeq     int    `xorm:"INTEGER NOT NULL"`
	WfpName    string `xorm:"VARCHAR(128)"`
	WfpDefault string `xorm:"VARCHAR(128) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}

/*
TableName xorm mapper
NewWorkflowParam struct map to table sys_workflow_param
*/
func (wfp NewWorkflowParam) TableName() string {
	return "sys_workflow_param"
}

// Save insert method
func (wfp NewWorkflowParam) Save() error {

	affected, err := utils.Engine.Insert(wfp)
	if err != nil {
		seelog.Errorf("utils.Engine.Insert Error : %v\n", err)
		return err
	}
	seelog.Debugf("%v insert : %v", affected, wfp)

	return nil

}

// Update method
func (wfp SysWorkflowParam) Update() error {

	affected, err := utils.Engine.ID(wfp.WfpId).Update(wfp)
	if err != nil {
		seelog.Errorf("utils.Engine.ID.Update Error : %v", err)
		return err
	}
	seelog.Debugf("%v update : %v", affected, wfp)

	return nil

}

/*
GetWorkflowParamByWfdID func(wfdid int) ([]SysWorkflowParam, error)
*/
func GetWorkflowParamByWfdID(wfdid int) ([]SysWorkflowParam, error) {

	params := make([]SysWorkflowParam, 0)

	if err := utils.Engine.Where("wfd_id = ?", wfdid).Asc("wfp_seq").Find(&params); err != nil {
		seelog.Errorf("utils.Engine.Where Error : %v", err)
		return nil, err
	}

	return params, nil

}

/*
DelWorkflowParamByID func(wfpid int)
*/
func DelWorkflowParamByID(wfpid int) error {

	wfp := new(SysWorkflowParam)
	wfp.WfpId = wfpid

	affected, err := utils.Engine.Delete(wfp)
	if err != nil {
		seelog.Errorf("utils.Engine.Delete Error : %v", err)
		return err
	}
	seelog.Debugf("%v delete : %v", affected, wfp)

	return nil
}

/*
GetLastWorkflowParamByWfdID func(wfdid int) (SysWorkflowParam, error)
*/
func GetLastWorkflowParamByWfdID(wfdid int) (SysWorkflowParam, error) {

	wfp := new(SysWorkflowParam)

	has, err := utils.Engine.Where("wfd_id = ?", wfdid).Desc("wfp_seq").Get(wfp)
	if err != nil {
		seelog.Errorf("utils.Engine.Desc.Get Error : %v", err)
		return SysWorkflowParam{}, err
	}

	if has {
		seelog.Debugf("Workflow Last Param : %v", wfp)
		return *wfp, nil
	}
	seelog.Debugf("This Dtl [%v] Has No Parameters", wfdid)

	return SysWorkflowParam{}, errors.New("No Records")

}
