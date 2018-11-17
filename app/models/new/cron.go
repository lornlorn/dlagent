package new

/*
SysCron struct map to table sys_cron
*/
type SysCron struct {
	CronId     int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	CronName   string `xorm:"VARCHAR(128) NOT NULL"`
	CronStatus string `xorm:"VARCHAR(8) NOT NULL"`
	CronTime   string `xorm:"VARCHAR(128) NOT NULL"`
	WfiId      int    `xorm:"INTEGER NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}
