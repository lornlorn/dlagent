package test

import (
	"app/models/new"
	"log"
)

/*
GetWFsTest func()
*/
func GetWFsTest() {
	workflows, err := new.GetWorkflows()
	if err != nil {
		log.Println(err)
	}
	log.Println(workflows)
}
