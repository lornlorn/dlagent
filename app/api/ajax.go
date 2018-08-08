package api

import (
	"app/models"
	"app/scheduler"
	"app/utils"
	"log"

	"github.com/tidwall/gjson"
)

// Ajax struct
type Ajax struct {
}

// StopScheduler (api API) func(a []byte) ([]byte, error)
func (ajax Ajax) StopScheduler(a []byte) []byte {
	log.Println(a)

	scheduler.Stop()

	ret := utils.GetAjaxRetJSON("0000", nil)

	return ret
}

/*
RunCMD func(data []byte) []byte
*/
func (ajax Ajax) RunCMD(data []byte) []byte {

	var retobj models.AjaxReturn

	shell := gjson.Get(string(data), "data.shell")
	cmd := gjson.Get(string(data), "data.cmd")

	result, err := scheduler.RunCmd(shell.String(), cmd.String())
	if err != nil {
		log.Printf("api.ajax.RunCMD scheduler.RunCmd Error : %v\n", err)
		retobj = utils.GetAjaxRetObj("9999", err)
	} else {
		retobj = utils.GetAjaxRetObj("0000", err)
	}
	log.Println(string(result))

	ret, _ := utils.Convert2JSON(retobj)

	return ret

}

/*
GetJobDtl func(data []byte) []byte
*/
func (ajax Ajax) GetJobDtl(data []byte) []byte {
	jobid := gjson.Get(string(data), "data.jobid")

	var retobj models.AjaxReturnWithData

	jobdtl, err := models.GetJobDtlByID(int(jobid.Int()))
	if err != nil {
		log.Printf("api.ajax.GetJobDtl models.GetJobByID Error : %v\n", err)
		retobj = utils.GetAjaxRetWithDataObj("9999", err, nil)
	} else {
		retobj = utils.GetAjaxRetWithDataObj("0000", err, jobdtl)
	}

	ret, _ := utils.Convert2JSON(retobj)

	return ret
}
