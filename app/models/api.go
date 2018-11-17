package models

import (
	"app/db"
	"log"
)

/*
SysApi struct map to table sys_api
*/
type SysApi struct {
	Key   string `xorm:"VARCHAR(32) NOT NULL"`
	Value string `xorm:"VARCHAR(32) NOT NULL"`
}

/*
GetApi func(apiType string,key string) (string, error)
Return reflect function map as string
*/
func GetApi(apiType string, key string) (string, error) {

	api := new(SysApi)
	// if err := db.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	has, err := db.Engine.Where("key = ?", apiType+"."+key).Get(api)
	if err != nil {
		log.Printf("models.api.GetApi -> db.Engine.Where Error : %v\n", err)
		return "", err
	}

	if has {
		log.Println(api.Value)
		return api.Value, nil
	}
	log.Println("models.api.GetApi : No Records")
	return "", nil

}
