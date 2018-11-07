package models

import (
	"app/db"
	"log"
)

/*
SysApiMap struct map table sys_api_map
*/
type SysApiMap struct {
	Type string `xorm:"VARCHAR(32) NOT NULL"`
	Key  string `xorm:"VARCHAR(32) NOT NULL"`
	Api  string `xorm:"VARCHAR(32) NOT NULL"`
}

/*
GetSysAPIMap func(apiType string,key string) (string, error)
Return reflect function map as string
*/
func GetSysAPIMap(apiType string, key string) (string, error) {

	sysApiMap := new(SysApiMap)
	// if err := db.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	has, err := db.Engine.Where("type = ? and key = ?", apiType, key).Get(sysApiMap)
	if err != nil {
		// return nil, err
		log.Println(err)
		return "", err
	}

	if has {
		log.Println(sysApiMap.Api)
		return sysApiMap.Api, nil
	}
	log.Println("models.sys_apimap.GetSysAPIMap : No Records")
	return "", nil

}
