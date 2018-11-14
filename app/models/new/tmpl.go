package new

/*
SysTmpl struct map to table sys_tmpl
*/
type SysTmpl struct {
	Key   string `xorm:"VARCHAR(32) NOT NULL"`
	Value string `xorm:"VARCHAR(128) NOT NULL"`
}
