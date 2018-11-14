package new

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
