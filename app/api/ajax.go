package api

import (
	"app/models"
	"app/scheduler"
	"app/utils"
	"log"
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
