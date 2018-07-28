package api

import (
	"app/models"
	"log"
)

/*
Tmpl struct
*/
type Tmpl struct {
}

/*
GetJobList func(data []byte) []byte
*/
func (tmpl Tmpl) GetJobList() interface{} {

	// var retobj models.AjaxReturn

	jobs, err := models.GetJobList()
	if err != nil {
		log.Printf("api.template.GetJobList models.GetJobList Error : %v", err)
		return nil
	}

	// ret, _ := utils.Convert2JSON(jobs)

	// return ret
	return jobs

	/*
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
	*/
}
