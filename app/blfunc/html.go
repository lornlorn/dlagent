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
GetCompWithParams func(reqBody []byte, reqURL url.Values) interface{}
*/
func (html HTML) GetCompWithParams(reqBody []byte, reqURL url.Values) interface{} {
	compid, err := strconv.Atoi(reqURL["CompId"][0])
	if err != nil {
		seelog.Errorf("strconv.Atoi Error : %v", err)
	}
	comp, err := models.GetComponentByID(compid)
	if err != nil {
		seelog.Errorf("models.GetComponentByID Error : %v", err)
	}
	params, err := models.GetParametersByCompID(compid)
	if err != nil {
		seelog.Errorf("models.GetParametersByCompID Error : %v", err)
	}

	var compWithParams models.ComponentWithParams
	compWithParams.Comp = comp
	compWithParams.Params = params

	seelog.Debugf("models.ComponentWithParams : %v", compWithParams)

	return compWithParams
}
