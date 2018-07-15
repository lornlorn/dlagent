package models

/*
Jobflow struct map table jobflow
*/
type Jobflow struct {
	JfId         int    `xorm:"INTEGER not null unique pk"`
	JfJobId      int    `xorm:"INTEGER not null"`
	JfName       string `xorm:"VARCHAR(128) not null"`
	JfSeq        int    `xorm:"INTEGER not null"`
	JfSh         string `xorm:"VARCHAR(64)"`
	JfCmd        string `xorm:"VARCHAR(256) not null"`
	JfStatus     string `xorm:"VARCHAR(16) not null"`
	JfRemark     string `xorm:"VARCHAR(512)"`
	JfCreate     string `xorm:"VARCHAR(32)"`
	JfCreatetime string `xorm:"VARCHAR(15)"`
	JfModify     string `xorm:"VARCHAR(32)"`
	JfModifytime string `xorm:"VARCHAR(15)"`
}
