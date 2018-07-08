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

	retobj := models.ReflectReturn{
		RetCode: "0000",
		RetMsg:  "成功",
		RetData: nil,
	}

	ret, _ := utils.Convert2JSON(retobj)

	return ret
}

/*
RunCMD func(data []byte) []byte
*/
func (api API) RunCMD(data []byte) []byte {

	var retobj = new(models.ReflectReturn)

	shell := gjson.Get(string(data), "data.shell")
	cmd := gjson.Get(string(data), "data.cmd")

	result, err := scheduler.RunCmd(shell.String(), cmd.String())
	if err != nil {
		log.Printf("scheduler.RunCmd Fail : %v\n", err)
		retobj = &models.ReflectReturn{
			RetCode: "9999",
			RetMsg:  utils.GetRetMsg("9999"),
			RetData: nil,
		}
	} else {
		retobj = &models.ReflectReturn{
			RetCode: "0000",
			RetMsg:  utils.GetRetMsg("0000"),
			RetData: nil,
		}
	}
	log.Println(result)

	ret, _ := utils.Convert2JSON(retobj)

	return ret

}
