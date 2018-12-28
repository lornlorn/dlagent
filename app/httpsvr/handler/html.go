package handler

import (
	"app/models"
	"app/utils"
	"html/template"
	"net/http"

	seelog "github.com/cihub/seelog"
	"github.com/gorilla/mux"
)

// HTMLHandler func(res http.ResponseWriter, req *http.Request)
func HTMLHandler(res http.ResponseWriter, req *http.Request) {

	seelog.Infof("Router HTML : %v", req.URL)
	key := mux.Vars(req)["key"]

	reqBody := utils.ReadRequestBody2JSON(req.Body)
	seelog.Debugf("Request Body : %v", string(reqBody))

	reqURL := req.URL.Query()
	seelog.Debugf("Request Params : %v", reqURL)

	tmplPages, err := models.GetTmpls(key)
	if err != nil {
		seelog.Errorf("models.GetTmpls Error : %v", err)
		return
	}
	seelog.Debugf("models.GetTmpls : %v", tmplPages)

	// tmpl, err := template.ParseFiles(fmt.Sprintf("views/test/%v.html", key))
	tmpl, err := template.ParseFiles(tmplPages...)
	if err != nil {
		seelog.Errorf("template.ParseFiles Error : %v", err)
		return
	}

	api, err := models.GetAPI("html", key)
	if err != nil {
		seelog.Errorf("models.GetAPI Error : %v", err)
		return
	}
	seelog.Debugf("models.GetAPI : %v", api)

	if api != "" {
		fc, err := utils.FuncCall(api, reqBody, reqURL)
		if err != nil {
			seelog.Errorf("utils.FuncCall Error : %v", err)
			return
		}
		seelog.Debugf("Return Data : %v", fc[0].Interface())

		tmpl.Execute(res, fc[0].Interface())
	} else {
		tmpl.Execute(res, nil)
	}

}
