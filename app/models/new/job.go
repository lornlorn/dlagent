package new

/*
SysJob struct map to table sys_job
*/
type SysJob struct {
	JobId     int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	JobUid    string `xorm:"VARCHAR(32) NOT NULL UNIQUE"`
	WfiId     int    `xorm:"INTEGER NOT NULL"`
	JobStatus string `xorm:"VARCHAR(8) NOT NULL"`
	JobSrc    string `xorm:"VARCHAR(8) NOT NULL"`
	JobLog    string `xorm:"VARCHAR(1024)"`
	StartTime string `xorm:"VARCHAR(15)"`
	EndTime   string `xorm:"VARCHAR(15)"`
}
