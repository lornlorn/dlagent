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
	has, err := utils.Engine.Where("key = ? and type = ?", key, apiType).Get(api)
	if err != nil {
		seelog.Errorf("utils.Engine.Where Error : %v", err)
		return "", err
	}

	if has {
		seelog.Debugf("API : %v", api.Value)
		return api.Value, nil
	}
	seelog.Debug("models.api.GetAPI : No Records")
	return "", nil

}
