package api

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
		return nil
	}

	compWithParams, err := models.GetComponentWithParamsByCompID(compid)
	if err != nil {
		seelog.Errorf("models.GetComponentWithParamsByCompID Error : %v", err)
		return nil
	}

	return compWithParams
}
