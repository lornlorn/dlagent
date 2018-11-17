package new

/*
SysPlan struct map to table sys_plan
*/
type SysPlan struct {
	PlanId     int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	PlanName   string `xorm:"VARCHAR(128) NOT NULL"`
	PlanStatus string `xorm:"VARCHAR(8) NOT NULL"`
	PlanTime   string `xorm:"VARCHAR(128) NOT NULL"`
	WfiId      int    `xorm:"INTEGER NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}
