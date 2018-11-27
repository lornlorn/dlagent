package api

import (
	"app/models"
	"app/utils"
	"log"
	"net/http"
)

/*
HTML struct
*/
type HTML struct {
}

/*
GetJobList func(req *http.Request) []byte
*/
func (html HTML) GetJobList(req *http.Request) interface{} {

	jobs, err := models.GetWorkflows()
	if err != nil {
		log.Printf("api.html.GetJobList -> models.GetWorkflows Error : %v", err)
		return nil
	}

	return jobs

}

/*
GetWorkflowAllByID func(req *http.Request) interface{}
*/
func (html HTML) GetWorkflowAllByID(req *http.Request) interface{} {
	wfiID := utils.GetParamFromRequest(req, "WfiId")
	return wfiID
}
