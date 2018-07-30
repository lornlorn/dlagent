package models

import (
	"app/db"
	"log"
)

/*
ApiMap struct map table api_map
*/
type ApiMap struct {
	Type string `xorm:"VARCHAR(32) NOT NULL"`
	Key  string `xorm:"VARCHAR(32) NOT NULL"`
	Api  string `xorm:"VARCHAR(32) NOT NULL"`
}

/*
GetAPIMap func(apiType string,key string) (string, error)
Return reflect function map as string
*/
func GetAPIMap(apiType string, key string) (string, error) {

	apiMap := new(ApiMap)
	// if err := db.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	has, err := db.Engine.Where("type = ? and key = ?", apiType, key).Get(apiMap)
	if err != nil {
		// return nil, err
		log.Println(err)
		return "", err
	}

	if has {
		log.Println(apiMap.Api)
		return apiMap.Api, nil
	}
	log.Println("models.apimap.GetAPIMap : No Records")
	return "", nil

}
