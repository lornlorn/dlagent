package models

import (
	"app/utils"
	"errors"

	seelog "github.com/cihub/seelog"
)

/*
SysWorkflowDtl struct map to table sys_workflow_dtl
*/
type SysWorkflowDtl struct {
	WfdId      int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	WfiId      int    `xorm:"INTEGER NOT NULL"`
	WfdSeq     int    `xorm:"INTEGER NOT NULL"`
	WfdName    string `xorm:"VARCHAR(128)"`
	WfdStatus  string `xorm:"VARCHAR(8) NOT NULL"`
	WfdShell   string `xorm:"VARCHAR(128)"`
	WfdCmd     string `xorm:"VARCHAR(1024) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(19)"`
	ModifyTime string `xorm:"VARCHAR(19)"`
}

/*
NewWorkflowDtl struct map to table sys_workflow_dtl
*/
type NewWorkflowDtl struct {
	WfiId      int    `xorm:"INTEGER NOT NULL"`
	WfdSeq     int    `xorm:"INTEGER NOT NULL"`
	WfdName    string `xorm:"VARCHAR(128)"`
	WfdStatus  string `xorm:"VARCHAR(8) NOT NULL"`
	WfdShell   string `xorm:"VARCHAR(128)"`
	WfdCmd     string `xorm:"VARCHAR(1024) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(19)"`
	ModifyTime string `xorm:"VARCHAR(19)"`
}

/*
TableName xorm mapper
NewWorkflowDtl struct map to table sys_workflow_dtl
*/
func (wfd NewWorkflowDtl) TableName() string {
	return "sys_workflow_dtl"
}

// Save insert method
func (wfd NewWorkflowDtl) Save() error {
	affected, err := utils.Engine.Insert(wfd)
	if err != nil {
		seelog.Errorf("utils.Engine.Insert Error : %v", err)
		return err
	}
	seelog.Debugf("%v insert : %v", affected, wfd)
	return nil
}

// Update method
func (wfd SysWorkflowDtl) Update() error {
	affected, err := utils.Engine.ID(wfd.WfdId).Update(wfd)
	if err != nil {
		seelog.Errorf("utils.Engine.ID.Update Error : %v", err)
		return err
	}
	seelog.Debugf("%v update : %v", affected, wfd)
	return nil
}

/*
GetWorkflowDtlByID func(wfdid int) (SysWorkflowDtl, error)
*/
func GetWorkflowDtlByID(wfdid int) (SysWorkflowDtl, error) {

	wfd := new(SysWorkflowDtl)
	wfd.WfdId = wfdid

	has, err := utils.Engine.Get(wfd)
	if err != nil {
		seelog.Errorf("utils.Engine.Get Error : %v", err)
		return SysWorkflowDtl{}, err
	}

	if !has {
		seelog.Debug("Get 0 row")
		return SysWorkflowDtl{}, errors.New("Get 0 rows")
	}

	seelog.Debugf("Workflow Detail : %v", wfd)

	return *wfd, nil
}

/*
GetWorkflowDtlByWfiID func(wfiid int) ([]SysWorkflowDtl, error)
*/
func GetWorkflowDtlByWfiID(wfiid int) ([]SysWorkflowDtl, error) {

	details := make([]SysWorkflowDtl, 0)

	if err := utils.Engine.Where("wfi_id = ?", wfiid).Asc("wfd_seq").Find(&details); err != nil {
		seelog.Errorf("utils.Engine.Where Error : %v", err)
		return nil, err
	}
	seelog.Debugf("Workflow Detail : %v", details)

	return details, nil
}

/*
DelWorkflowDtlByID func(wfdid int)
*/
func DelWorkflowDtlByID(wfdid int) error {
	wfd := new(SysWorkflowDtl)
	wfd.WfdId = wfdid
	affected, err := utils.Engine.Delete(wfd)
	if err != nil {
		seelog.Errorf("utils.Engine.Delete Error : %v", err)
		return err
	}
	seelog.Debugf("%v delete : %v", affected, wfd)

	return nil
}
