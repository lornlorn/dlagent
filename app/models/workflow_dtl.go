package models

import (
	"app/utils"
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
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
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
