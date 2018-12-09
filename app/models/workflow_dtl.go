package models

import (
	"app/utils"
	"errors"
	"log"
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
	// affected, err := utils.Engine.Insert(d)
	_, err := utils.Engine.Insert(wfd)
	if err != nil {
		log.Printf("models.workflow_dtl.Save -> utils.Engine.Insert Error : %v\n", err)
		return err
	}
	return nil
}

// Update method
func (wfd SysWorkflowDtl) Update() error {
	// affected, err := utils.Engine.Insert(d)
	_, err := utils.Engine.ID(wfd.WfdId).Update(wfd)
	if err != nil {
		log.Printf("models.workflow_dtl.Update -> utils.Engine.ID.Update Error : %v\n", err)
		return err
	}
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
		log.Printf("models.workflow_dtl.GetWorkflowByID -> utils.Engine.Get Error : %v\n", err)
		return SysWorkflowDtl{}, err
	}

	if !has {
		return SysWorkflowDtl{}, errors.New("Get 0 rows")
	}

	log.Println(wfd)

	return *wfd, nil
}

/*
GetWorkflowDtlByWfiID func(wfiid int) ([]SysWorkflowDtl, error)
*/
func GetWorkflowDtlByWfiID(wfiid int) ([]SysWorkflowDtl, error) {

	details := make([]SysWorkflowDtl, 0)

	if err := utils.Engine.Where("wfi_id = ?", wfiid).Asc("wfd_seq").Find(&details); err != nil {
		log.Printf("models.workflow_dtl.GetWorkflowDtlByWfiID -> utils.Engine.Where Error : %v\n", err)
		return nil, err
	}

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
		log.Printf("models.workflow_dtl.DelWorkflowDtlByID -> utils.Engine.Delete Error : %v\n", err)
		return err
	}
	log.Printf("删除%v条记录\n", affected)
	return nil
}
