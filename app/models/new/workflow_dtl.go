package new

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
