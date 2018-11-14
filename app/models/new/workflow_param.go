package new

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
