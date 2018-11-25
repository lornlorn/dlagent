package test

import (
	"app/models"
	"app/utils"
	"log"

	"github.com/tidwall/gjson"
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

/*
DelWfByID func()
*/
func DelWfByID(data []byte) []byte {
	wfiid := gjson.Get(string(data), "data.WfiId")
	err := models.DelWorkflowByID(int(wfiid.Int()))
	if err != nil {
		log.Printf("test.test.DelWfByID -> models.DelWorkflowByID Error : %v\n", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}
