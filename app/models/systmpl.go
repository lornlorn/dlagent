package models

import (
	"app/utils"
	"errors"
)

/*
SysTmpl struct map to table sys_tmpl
*/
type SysTmpl struct {
	Key   string `xorm:"VARCHAR(32) NOT NULL"`
	Value string `xorm:"VARCHAR(128) NOT NULL"`
	Seq   int    `xorm:"INTEGER NOT NULL"`
}

/*
GetTmpls func(key string) ([]string, error)
Return HTML pages path as []string
*/
func GetTmpls(key string) ([]string, error) {
	tmpls := make([]SysTmpl, 0)

	if err := utils.Engine.Where("key = ?", key).Asc("seq").Find(&tmpls); err != nil {
		// seelog.Errorf("utils.Engine.Where Error : %v", err)
		return nil, err
	}

	if len(tmpls) == 0 {
		// seelog.Debug("models.tmpl.GetTmpls : No Records")
		return nil, errors.New("models.tmpl.GetTmpls : No Records")
	}

	pages := make([]string, len(tmpls))
	for i, v := range tmpls {
		pages[i] = v.Value
	}

	return pages, nil
}
