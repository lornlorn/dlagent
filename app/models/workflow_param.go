package models

import (
	"app/utils"
	"log"
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
GetWorkflowParamByWfdID func(wfdid int) ([]SysWorkflowParam, error)
*/
func GetWorkflowParamByWfdID(wfdid int) ([]SysWorkflowParam, error) {

	params := make([]SysWorkflowParam, 0)

	if err := utils.Engine.Where("wfd_id = ?", wfdid).Asc("wfp_seq").Find(&params); err != nil {
		log.Printf("models.workflow_dtl.GetWorkflowParamByWfdID -> utils.Engine.Where Error : %v\n", err)
		return nil, err
	}

	return params, nil
}
