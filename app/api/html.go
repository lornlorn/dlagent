package api

import (
	"app/models"
	"log"
)

/*
HTML struct
*/
type HTML struct {
}

/*
GetJobList func(data []byte) []byte
*/
func (html HTML) GetJobList() interface{} {

	jobs, err := models.GetWorkflows()
	if err != nil {
		log.Printf("api.html.GetJobList -> models.GetWorkflows Error : %v", err)
		return nil
	}

	return jobs

}
