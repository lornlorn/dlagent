package blfunc

import (
	"app/models"
	"net/url"
	"strconv"

	seelog "github.com/cihub/seelog"
)

/*
HTML struct
*/
type HTML struct {
}

/*
GetWorkflowInfDtl func(reqBody []byte, reqURL url.Values) interface{}
*/
func (html HTML) GetWorkflowInfDtl(reqBody []byte, reqURL url.Values) interface{} {
	wfiid, err := strconv.Atoi(reqURL["WfiId"][0])
	if err != nil {
		seelog.Errorf("strconv.Atoi Error : %v", err)
	}
	wfi, err := models.GetWorkflowByID(wfiid)
	if err != nil {
		seelog.Errorf("models.GetWorkflowByID Error : %v", err)
	}
	seelog.Debugf("models.GetWorkflowByID : %v", wfi)
	return wfi
}

/*
GetWorkflowDtlParam func(reqBody []byte, reqURL url.Values) interface{}
*/
func (html HTML) GetWorkflowDtlParam(reqBody []byte, reqURL url.Values) interface{} {
	wfdid, err := strconv.Atoi(reqURL["WfdId"][0])
	if err != nil {
		seelog.Errorf("strconv.Atoi Error : %v", err)
	}
	wfd, err := models.GetWorkflowDtlByID(wfdid)
	if err != nil {
		seelog.Errorf("models.GetWorkflowDtlByID Error : %v", err)
	}
	wfp, err := models.GetWorkflowParamByWfdID(wfdid)
	if err != nil {
		seelog.Errorf("models.GetWorkflowParamByWfdID Error : %v", err)
	}

	var dtlAndParam models.WorkflowDtlWithParams
	dtlAndParam.WFD = wfd
	dtlAndParam.WFP = wfp

	seelog.Debugf("models.WorkflowDtlWithParams : %v", dtlAndParam)

	return dtlAndParam
}
