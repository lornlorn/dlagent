package models

import (
	"app/db"
	"errors"
	"log"
)

/*
SysTmplMap struct map table sys_tmpl_map
*/
type SysTmplMap struct {
	Key  string `xorm:"VARCHAR(32) NOT NULL"`
	Page string `xorm:"VARCHAR(128) NOT NULL"`
}

/*
GetSysTmplMap func(key string) ([]string, error)
Return HTML pages path as []string
*/
func GetSysTmplMap(key string) ([]string, error) {

	sysTmplMap := make([]SysTmplMap, 0)
	// if err := db.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	if err := db.Engine.Where("key = ?", key).Find(&sysTmplMap); err != nil {
		// return nil, err
		log.Println(err)
		return nil, err
	}

	if len(sysTmplMap) == 0 {
		return nil, errors.New("models.sys_tmplmap.GetSysTmplMap : No Records")
	}

	pages := make([]string, len(sysTmplMap))
	for i, v := range sysTmplMap {
		log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
		pages[i] = v.Page
	}
	return pages, nil
}
