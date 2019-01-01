package new

/*
SysParameter struct map to table sys_parameter
*/
type SysParameter struct {
	Key   string `xorm:"VARCHAR(32) NOT NULL UNIQUE"`
	Value string `xorm:"VARCHAR(32) NOT NULL"`
}
