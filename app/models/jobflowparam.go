package models

/*
JobflowParam struct map table jobflow_param
*/
type JobflowParam struct {
	JfpId         int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	JfpJfId       int    `xorm:"INTEGER NOT NULL"`
	JfpSeq        int    `xorm:"INTEGER NOT NULL"`
	JfpParameter  string `xorm:"VARCHAR(64) NOT NULL"`
	JfpDefault    string `xorm:"VARCHAR(512)"`
	JfpStatus     string `xorm:"VARCHAR(16) NOT NULL"`
	JfpRemark     string `xorm:"VARCHAR(512)"`
	JfpCreate     string `xorm:"VARCHAR(32)"`
	JfpCreatetime string `xorm:"VARCHAR(15)"`
	JfpModify     string `xorm:"VARCHAR(32)"`
	JfpModifytime string `xorm:"VARCHAR(15)"`
}
