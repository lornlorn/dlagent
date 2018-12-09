package api

import (
	"app/models"
	"app/utils"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"
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
AddWorkflow func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) AddWorkflow(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	wfiname := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiName")
	wfistatus := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiStatus")
	wfidesc := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiDesc")
	var wfi models.NewWorkflowInf
	wfi = models.NewWorkflowInf{
		WfiName:    wfiname.String(),
		WfiStatus:  wfistatus.String(),
		WfiDesc:    wfidesc.String(),
		CreateTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	log.Println(wfi)
	err := wfi.Save()
	if err != nil {
		log.Printf("api.ajax.AddWorkflow -> wfi.Save Error : %v\n", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
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
AddWorkflowDtl func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) AddWorkflowDtl(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	wfiid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiId")
	wfdSeq := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdSeq")
	wfdname := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdName")
	wfdstatus := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdStatus")
	wfdshell := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdShell")
	wfdcmd := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdCmd")
	var wfd models.NewWorkflowDtl
	wfd = models.NewWorkflowDtl{
		WfiId:      int(wfiid.Int()),
		WfdSeq:     int(wfdSeq.Int()),
		WfdName:    wfdname.String(),
		WfdStatus:  wfdstatus.String(),
		WfdShell:   wfdshell.String(),
		WfdCmd:     wfdcmd.String(),
		CreateTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	log.Println(wfd)
	err := wfd.Save()
	if err != nil {
		log.Printf("api.ajax.AddWorkflowDtl -> wfd.Save Error : %v\n", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
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
