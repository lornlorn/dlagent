package handler

import (
	"app/models"
	"app/utils"
	"errors"
	"net/http"

	seelog "github.com/cihub/seelog"
	"github.com/gorilla/mux"
)

// AjaxHandler func(res http.ResponseWriter, req *http.Request)
func AjaxHandler(res http.ResponseWriter, req *http.Request) {
	seelog.Infof("Router Ajax : %v", req.URL)
	key := mux.Vars(req)["key"]

	reqBody := utils.ReadRequestBody2JSON(req.Body)
	seelog.Debugf("Request Body : %v", string(reqBody))

	reqURL := req.URL.Query()
	seelog.Debugf("Request Params : %v", reqURL)

	api, err := models.GetAPI("ajax", key)
	if err != nil {
		seelog.Errorf("models.GetAPI Error : %v", err)
		res.Write(utils.GetAjaxRetJSON("9999", errors.New("获取API失败")))
		return
	}
	seelog.Debugf("models.GetAPI : %v", api)

	if api != "" {
		fc, err := utils.FuncCall(api, reqBody, reqURL)
		if err != nil {
			seelog.Errorf("utils.FuncCall Error : %v", err)
			res.Write(utils.GetAjaxRetJSON("9999", errors.New("API调用失败")))
			return
		}
		seelog.Debugf("Return Data : %v", string(fc[0].Bytes()))

		res.Write(fc[0].Bytes())
	} else {
		res.Write(utils.GetAjaxRetJSON("9999", errors.New("非法API")))
	}

}
