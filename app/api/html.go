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

	jobs, err := models.GetJobs("tool")
	if err != nil {
		log.Printf("api.html.GetJobList -> models.GetJobs Error : %v", err)
		return nil
	}

	return jobs

}
