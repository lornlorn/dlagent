package api

import (
	"app/models"
	"log"
)

/*
Tmpl struct
*/
type Tmpl struct {
}

/*
GetJobList func(data []byte) []byte
*/
func (tmpl Tmpl) GetJobList() interface{} {

	jobs, err := models.GetJobs("tool")
	if err != nil {
		log.Printf("api.template.GetJobList models.GetJobs Error : %v", err)
		return nil
	}

	return jobs

}
