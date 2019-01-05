package blfunc


/*
GetWorkflows func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) GetWorkflows(reqBody []byte, reqURL url.Values) []byte {
	workflows, err := models.GetWorkflows()
	if err != nil {
		seelog.Errorf("models.GetWorkflows Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	seelog.Debugf("models.GetWorkflows : %v", workflows)
	return utils.Convert2JSON(workflows)
}

/*
AddWorkflowInf func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) AddWorkflowInf(reqBody []byte, reqURL url.Values) []byte {
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
	seelog.Debugf("models.NewWorkflowInf : %v", wfi)
	err := wfi.Save()
	if err != nil {
		seelog.Errorf("wfi.Save Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
DelWorkflowInf func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) DelWorkflowInf(reqBody []byte, reqURL url.Values) []byte {
	wfiid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiId")
	err := models.DelWorkflowByID(int(wfiid.Int()))
	if err != nil {
		seelog.Errorf("models.DelWorkflowByID Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
UpdateWorkflowInf func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) UpdateWorkflowInf(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	wfiid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiId")
	wfiname := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiName")
	wfistatus := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiStatus")
	wfidesc := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiDesc")
	var wfi models.SysWorkflowInf
	wfi = models.SysWorkflowInf{
		WfiId:      int(wfiid.Int()),
		WfiName:    wfiname.String(),
		WfiStatus:  wfistatus.String(),
		WfiDesc:    wfidesc.String(),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	seelog.Debugf("models.SysWorkflowInf : %v", wfi)
	err := wfi.Update()
	if err != nil {
		seelog.Errorf("wfi.Update Error : %v", err)
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
		seelog.Errorf("strconv.Atoi Error : %v\n", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	wfds, err := models.GetWorkflowDtlByWfiID(wfiid)
	if err != nil {
		seelog.Errorf("models.GetWorkflowDtlByWfiID Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	seelog.Debugf("models.GetWorkflowDtlByWfiID : %v", wfds)

	return utils.Convert2JSON(wfds)
}

/*
AddWorkflowDtl func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) AddWorkflowDtl(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	wfiid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfiId")
	wfdseq := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdSeq")
	wfdname := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdName")
	wfdstatus := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdStatus")
	wfdshell := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdShell")
	wfdcmd := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdCmd")
	var wfd models.NewWorkflowDtl
	wfd = models.NewWorkflowDtl{
		WfiId:      int(wfiid.Int()),
		WfdSeq:     int(wfdseq.Int()),
		WfdName:    wfdname.String(),
		WfdStatus:  wfdstatus.String(),
		WfdShell:   wfdshell.String(),
		WfdCmd:     wfdcmd.String(),
		CreateTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	seelog.Debugf("models.NewWorkflowDtl : %v", wfd)
	err := wfd.Save()
	if err != nil {
		seelog.Errorf("wfd.Save Error : %v", err)
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
		seelog.Errorf("models.DelWorkflowDtlByID Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
UpdateWorkflowDtl func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) UpdateWorkflowDtl(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	wfdid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdId")
	wfdseq := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdSeq")
	wfdname := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdName")
	wfdstatus := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdStatus")
	wfdshell := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdShell")
	wfdcmd := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdCmd")
	var wfd models.SysWorkflowDtl
	wfd = models.SysWorkflowDtl{
		WfdId:      int(wfdid.Int()),
		WfdSeq:     int(wfdseq.Int()),
		WfdName:    wfdname.String(),
		WfdStatus:  wfdstatus.String(),
		WfdShell:   wfdshell.String(),
		WfdCmd:     wfdcmd.String(),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	seelog.Debugf("models.SysWorkflowDtl : %v", wfd)
	err := wfd.Update()
	if err != nil {
		seelog.Errorf("wfd.Update Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
AddWorkflowParam func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) AddWorkflowParam(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板
	wfdid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdId")
	wfpseq := utils.GetJSONResultFromRequestBody(reqBody, "data.WfpSeq")
	wfpname := utils.GetJSONResultFromRequestBody(reqBody, "data.WfpName")
	wfpdefault := utils.GetJSONResultFromRequestBody(reqBody, "data.WfpDefault")
	var wfp models.NewWorkflowParam
	wfp = models.NewWorkflowParam{
		WfdId:      int(wfdid.Int()),
		WfpSeq:     int(wfpseq.Int()),
		WfpName:    wfpname.String(),
		WfpDefault: wfpdefault.String(),
		CreateTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
	}
	seelog.Debugf("models.NewWorkflowParam : %v", wfp)
	err := wfp.Save()
	if err != nil {
		seelog.Errorf("wfp.Save Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
DelWorkflowParam func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) DelWorkflowParam(reqBody []byte, reqURL url.Values) []byte {
	wfdid := utils.GetJSONResultFromRequestBody(reqBody, "data.WfdId")

	lastwfp, err := models.GetLastWorkflowParamByWfdID(int(wfdid.Int()))
	if err != nil {
		seelog.Errorf("models.GetLastWorkflowParamByWfdID Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}

	err = models.DelWorkflowParamByID(lastwfp.WfpId)
	if err != nil {
		seelog.Errorf("models.DelWorkflowParamByID Error : %v", err)
		return utils.GetAjaxRetJSON("9999", nil)
	}
	return utils.GetAjaxRetJSON("0000", nil)
}

/*
UpdateWorkflowParam func(reqBody []byte, reqURL url.Values) []byte
*/
func (ajax Ajax) UpdateWorkflowParam(reqBody []byte, reqURL url.Values) []byte {
	nowTime := time.Now()
	timeFormat := "2006-01-02 15:04:05" // 时间格式化模板

	params := utils.ReadJSONData2Array(reqBody, "data.paramlist")
	seelog.Debugf("Parameters : %v", params)

	for _, v := range params {
		wfp := models.SysWorkflowParam{
			WfpId:      int(v.Get("WfpId").Int()),
			WfpName:    v.Get("WfpName").String(),
			WfpDefault: v.Get("WfpDefault").String(),
			ModifyTime: fmt.Sprintf("%v", nowTime.Format(timeFormat)),
		}
		err := wfp.Update()
		if err != nil {
			seelog.Errorf("wfp.Update Error : %v", err)
			return utils.GetAjaxRetJSON("9999", nil)
		}
	}
	return utils.GetAjaxRetJSON("0000", nil)
}