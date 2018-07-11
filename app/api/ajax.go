package api

import (
	"app/models"
	"app/scheduler"
	"app/utils"
	"log"

	"github.com/tidwall/gjson"
)

// API struct
type API struct {
}

// StopScheduler (api API) func(a []byte) ([]byte, error)
func (api API) StopScheduler(a []byte) []byte {
	log.Println(a)

	scheduler.Stop()

	ret := utils.GetAjaxRetJSON("0000", nil)

	return ret
}

/*
RunCMD func(data []byte) []byte
*/
func (api API) RunCMD(data []byte) []byte {

	var retobj models.AjaxReturn

	shell := gjson.Get(string(data), "data.shell")
	cmd := gjson.Get(string(data), "data.cmd")

	result, err := scheduler.RunCmd(shell.String(), cmd.String())
	if err != nil {
		log.Printf("scheduler.RunCmd Fail : %v\n", err)
		retobj = utils.GetAjaxRetObj("9999", err)
	} else {
		retobj = utils.GetAjaxRetObj("0000", err)
	}
	log.Println(string(result))

	ret, _ := utils.Convert2JSON(retobj)

	return ret

}
