package api

import (
	"app/models"
	"app/utils"
	"log"
	"net/url"
)

/*
HTML struct
*/
type HTML struct {
}

/*
GetJobList func(reqBody []byte, reqURL url.Values) []byte
*/
func (html HTML) GetJobList(reqBody []byte, reqURL url.Values) interface{} {

	jobs, err := models.GetWorkflows()
	if err != nil {
		log.Printf("api.html.GetJobList -> models.GetWorkflows Error : %v", err)
		return nil
	}

	return jobs

}

/*
GetWorkflowAllByID func(reqBody []byte, reqURL url.Values) interface{}
*/
func (html HTML) GetWorkflowAllByID(reqBody []byte, reqURL url.Values) interface{} {
	wfiID := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiId")
	return wfiID
}
