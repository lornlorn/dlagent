package models

import (
	"app/utils"

	seelog "github.com/cihub/seelog"
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
		seelog.Errorf("utils.Engine.Where Error : %v", err)
		return "", err
	}

	if has {
		// log.Println(api.Value)
		seelog.Debugf("API : %v", api.Value)
		return api.Value, nil
	}
	seelog.Debug("models.api.GetAPI : No Records")
	return "", nil

}
