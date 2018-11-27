package models

import (
	"app/utils"
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
GetAPI func(apiType string,key string) (string, error)
Return reflect function map as string
*/
func GetAPI(apiType string, key string) (string, error) {

	api := new(SysApi)
	// if err := utils.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	has, err := utils.Engine.Where("key = ? and type = ?", key, apiType).Get(api)
	if err != nil {
		log.Printf("models.api.GetAPI -> utils.Engine.Where Error : %v\n", err)
		return "", err
	}

	if has {
		// log.Println(api.Value)
		return api.Value, nil
	}
	log.Println("models.api.GetAPI : No Records")
	return "", nil

}
