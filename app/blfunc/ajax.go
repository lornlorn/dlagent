package blfunc

import (
	"app/models"
	"app/scheduler"
	"app/utils"
	"fmt"
	"net/url"
	"time"

	seelog "github.com/cihub/seelog"
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
RunCMD func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) RunCMD(reqBody []byte, reqURL url.Values) []byte {
	wfiid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiId")
	wfiName := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiName")
	seelog.Debugf("Workflow ID : [%v]", wfiid)

	command := "F:\\WorkSpace\\Proj\\GoProj\\dlagent\\src\\data\\cmd\\test.bat"
	envs := []string{"TEST_ARG=xxx"}
	params := []string{"AAA", wfiName.String(), "哈哈哈"}
	ret, err := scheduler.Run(command, envs, params...)
	if err != nil {
		seelog.Errorf("scheduler.Run Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}
	result := string(ret)
	seelog.Debugf("Command result : %v", result)

	return utils.GetAjaxRetWithDataJSON("0000", nil, result)

}

/*
GetComponents func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) GetComponents(reqBody []byte, reqURL url.Values) []byte {
	components, err := models.GetComponents()
	if err != nil {
		seelog.Errorf("models.GetComponents Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}
	seelog.Debugf("models.GetComponents : %v", components)
	return utils.Convert2JSON(components)
}

/*
DelComponentByID func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) DelComponentByID(reqBody []byte, reqURL url.Values) []byte {
	compid := utils.GetJSONResultFromRequestBody(reqBody, "data.CompId")
	err := models.DelComponentByID(int(compid.Int()))
	if err != nil {
		seelog.Errorf("models.DelComponentByID Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
NewComponent func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) NewComponent(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	compname := utils.GetJSONResultFromRequestBody(reqBody, "data.CompName")
	compcmd := utils.GetJSONResultFromRequestBody(reqBody, "data.CompCmd")
	var comp models.NewComponent
	comp = models.NewComponent{
		CompName:   compname.String(),
		CompCmd:    compcmd.String(),
		CreateTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	seelog.Debugf("models.NewComponent : %v", comp)
	err := comp.Save()
	if err != nil {
		seelog.Errorf("comp.Save Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
NewParameter func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) NewParameter(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	compid := utils.GetJSONResultFromRequestBody(reqBody, "data.CompId")
	paramseq := utils.GetJSONResultFromRequestBody(reqBody, "data.ParamSeq")
	paramname := utils.GetJSONResultFromRequestBody(reqBody, "data.ParamName")
	paramdefault := utils.GetJSONResultFromRequestBody(reqBody, "data.ParamDefault")
	var param models.NewParameter
	param = models.NewParameter{
		CompId:       int(compid.Int()),
		ParamSeq:     int(paramseq.Int()),
		ParamName:    paramname.String(),
		ParamDefault: paramdefault.String(),
		CreateTime:   fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		ModifyTime:   fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	seelog.Debugf("models.NewParameter : %v", param)
	err := param.Save()
	if err != nil {
		seelog.Errorf("param.Save Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
DelLastParameterByCompID func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) DelLastParameterByCompID(reqBody []byte, reqURL url.Values) []byte {
	compid := utils.GetJSONResultFromRequestBody(reqBody, "data.CompId")

	lastparam, err := models.GetLastParameterByCompID(int(compid.Int()))
	if err != nil {
		seelog.Errorf("models.GetLastParameterByCompID Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}

	err = models.DelParameterByID(lastparam.ParamId)
	if err != nil {
		seelog.Errorf("models.DelParameterByID Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
UpdateParameters func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) UpdateParameters(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板

	params := utils.ReadJSONData2Array(reqBody, "data.paramlist")
	seelog.Debugf("Parameters : %v", params)

	for i, v := range params {
		param := models.TbParameter{
			ParamId:      int(v.Get("ParamId").Int()),
			CompId:       int(v.Get("CompId").Int()),
			ParamSeq:     i + 1,
			ParamName:    v.Get("ParamName").String(),
			ParamDefault: v.Get("ParamDefault").String(),
			ModifyTime:   fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		}
		err := param.Update()
		if err != nil {
			seelog.Errorf("param.Update Error : %v", err)
			return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
		}
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
UpdateComponent func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) UpdateComponent(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	compid := utils.GetJSONResultFromRequestBody(reqBody, "data.CompId")
	compname := utils.GetJSONResultFromRequestBody(reqBody, "data.CompName")
	compcmd := utils.GetJSONResultFromRequestBody(reqBody, "data.CompCmd")
	var comp models.TbComponent
	comp = models.TbComponent{
		CompId:     int(compid.Int()),
		CompName:   compname.String(),
		CompCmd:    compcmd.String(),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	seelog.Debugf("models.TbComponent : %v", comp)
	err := comp.Update()
	if err != nil {
		seelog.Errorf("comp.Update Error : %v", err)
		return utils.GetAjaxRetWithDataJSON("9999", nil, err.Error())
	}
	return utils.GetAjaxRetJSON("0000", nil)
}
