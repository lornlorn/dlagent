package models

import (
	"app/utils"
	"errors"
	"log"
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
	// if err := utils.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	if err := utils.Engine.Where("key = ?", key).Asc("seq").Find(&tmpls); err != nil {
		// return nil, err
		log.Printf("models.tmpl.GetTmpls -> utils.Engine.Where Error : %v\n", err)
		return nil, err
	}

	if len(tmpls) == 0 {
		return nil, errors.New("models.tmpl.GetTmpls : No Records")
	}

	pages := make([]string, len(tmpls))
	for i, v := range tmpls {
		log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
		pages[i] = v.Value
	}
	return pages, nil
}
