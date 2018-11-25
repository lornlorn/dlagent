package test

import (
	"app/models"
	"app/utils"
	"log"
)

/*
GetWFsTest func()
*/
func GetWFsTest() []byte {
	workflows, err := models.GetWorkflows()
	if err != nil {
		log.Println(err)
	}
	log.Println(workflows)
	return utils.Convert2JSON(workflows)
}
