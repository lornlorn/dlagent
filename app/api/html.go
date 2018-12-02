package api

import (
	"app/models"
	"app/utils"
	"log"
	"net/url"
	"strconv"
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

/*
GetWorkflowInfDtl func(reqBody []byte, reqURL url.Values) interface{}
*/
func (html HTML) GetWorkflowInfDtl(reqBody []byte, reqURL url.Values) interface{} {
	wfiid, err := strconv.Atoi(reqURL["WfiId"][0])
	if err != nil {
		log.Printf("api.html.GetWorkflowInfDtl -> strconv.Atoi Error : %v\n", err)
	}
	wfi, err := models.GetWorkflowByID(wfiid)
	if err != nil {
		log.Printf("api.html.GetWorkflowInfDtl -> models.GetWorkflowByID Error : %v\n", err)
	}
	log.Println("---")
	log.Println(wfi)
	return wfi
}

/*
GetWorkflowDtlParam func(reqBody []byte, reqURL url.Values) interface{}
*/
func (html HTML) GetWorkflowDtlParam(reqBody []byte, reqURL url.Values) interface{} {
	wfdid, err := strconv.Atoi(reqURL["WfdId"][0])
	if err != nil {
		log.Printf("api.html.GetWorkflowDtlParam -> strconv.Atoi Error : %v\n", err)
	}
	wfd, err := models.GetWorkflowDtlByID(wfdid)
	if err != nil {
		log.Printf("api.html.GetWorkflowDtlParam -> models.GetWorkflowDtlByID Error : %v\n", err)
	}
	wfp, err := models.GetWorkflowParamByWfdID(wfdid)
	if err != nil {
		log.Printf("api.html.GetWorkflowDtlParam -> models.GetWorkflowParamByWfdID Error : %v\n", err)
	}

	var dtlAndParam models.WorkflowDtlWithParams
	dtlAndParam.WFD = wfd
	dtlAndParam.WFP = wfp

	log.Println("===")
	log.Println(dtlAndParam)

	return dtlAndParam
}
