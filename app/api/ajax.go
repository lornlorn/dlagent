package api

import (
	"app/models"
	"app/utils"
	"log"
	"net/url"
	"strconv"
)

// Ajax struct
type Ajax struct {
}

// StopScheduler (api API) func(a []byte) ([]byte, error)
/*
func (ajax Ajax) StopScheduler(a []byte) []byte {
	log.Println(a)

	scheduler.Stop()

	ret := utils.GetAjaxRetJSON("0000", nil)

	return ret
}
*/

/*
RunCMD func(req *http.Request) []byte
*/
/*
func (ajax Ajax) RunCMD(req *http.Request) []byte {

	var retobj utils.AjaxReturn

	shell := utils.GetJSONResultFromRequest(req, "data.shell")
	cmd := utils.GetJSONResultFromRequest(req, "data.cmd")

	result, err := scheduler.RunCmd(shell.String(), cmd.String())
	if err != nil {
		log.Printf("api.ajax.RunCMD scheduler.RunCmd Error : %v\n", err)
		retobj = utils.GetAjaxRetObj("9999", err)
	} else {
		retobj = utils.GetAjaxRetObj("0000", err)
	}
	log.Println(string(result))

	ret := utils.Convert2JSON(retobj)

	return ret

}
*/

/*
GetJobDtl func(reqBody []byte, reqURL url.Values) []byte
暂时没用
*/
func (ajax Ajax) GetJobDtl(reqBody []byte, reqURL url.Values) []byte {

	jobid := utils.GetJSONResultFromRequestBody(reqBody, "data.jobid")

	var retobj utils.AjaxReturnWithData
	jobdtl, err := models.GetWorkflowsAllByID(int(jobid.Int()))
	if err != nil {
		log.Printf("api.ajax.GetJobDtl -> models.GetJobDtlByID Error : %v\n", err)
		retobj = utils.GetAjaxRetWithDataObj("9999", err, nil)
	} else {
		retobj = utils.GetAjaxRetWithDataObj("0000", err, jobdtl)
	}

	return utils.Convert2JSON(retobj)
}

/*
GetWorkflows func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) GetWorkflows(reqBody []byte, reqURL url.Values) []byte {
	workflows, err := models.GetWorkflows()
	if err != nil {
		log.Printf("api.ajax.GetWorkflows ->  models.GetWorkflows Error : %v\n", err)
	}
	log.Println(workflows)
	return utils.Convert2JSON(workflows)
}

/*
DelWorkflow func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) DelWorkflow(reqBody []byte, reqURL url.Values) []byte {
	wfiid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiId")
	err := models.DelWorkflowByID(int(wfiid.Int()))
	if err != nil {
		log.Printf("api.ajax.DelWorkflow -> models.DelWorkflowByID Error : %v\n", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
GetWorkflowDtl func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) GetWorkflowDtl(reqBody []byte, reqURL url.Values) []byte {
	wfiid, err := strconv.Atoi(reqURL["WfiId"][0])
	if err != nil {
		log.Printf("api.ajax.GetWorkflowDtl -> strconv.Atoi Error : %v\n", err)
	}
	wfds, err := models.GetWorkflowDtlByWfiID(wfiid)
	if err != nil {
		log.Printf("api.ajax.GetWorkflowDtl -> models.GetWorkflowDtlByWfiID Error : %v\n", err)
	}
	log.Println(wfds)

	return utils.Convert2JSON(wfds)
}

/*
DelWorkflowDtl func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) DelWorkflowDtl(reqBody []byte, reqURL url.Values) []byte {
	wfdid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdId")
	err := models.DelWorkflowDtlByID(int(wfdid.Int()))
	if err != nil {
		log.Printf("api.ajax.DelWorkflowDtl -> models.DelWorkflowDtlByID Error : %v\n", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}
