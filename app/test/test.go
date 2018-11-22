package test

import (
	"app/models"
	"log"
)

/*
GetWFsTest func()
*/
func GetWFsTest() {
	workflows, err := models.GetWorkflows()
	if err != nil {
		log.Println(err)
	}
	log.Println(workflows)
}
