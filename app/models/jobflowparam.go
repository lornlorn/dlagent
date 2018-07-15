package models

/*
JobflowParam struct map table jobflow_param
*/
type JobflowParam struct {
	JfpId         int    `xorm:"INTEGER not null unique pk"`
	JfpJfId       int    `xorm:"INTEGER not null"`
	JfpSeq        int    `xorm:"INTEGER not null"`
	JfpParameter  string `xorm:"VARCHAR(64) not null"`
	JfpDefault    string `xorm:"VARCHAR(512)"`
	JfpStatus     string `xorm:"VARCHAR(16) not null"`
	JfpRemark     string `xorm:"VARCHAR(512)"`
	JfpCreate     string `xorm:"VARCHAR(32)"`
	JfpCreatetime string `xorm:"VARCHAR(15)"`
	JfpModify     string `xorm:"VARCHAR(32)"`
	JfpModifytime string `xorm:"VARCHAR(15)"`
}
