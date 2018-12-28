package handler

import (
	"app/utils"
	"html/template"
	"net/http"

	seelog "github.com/cihub/seelog"
	"github.com/gorilla/mux"
)

// TestHandler func(res http.ResponseWriter, req *http.Request)
func TestHandler(res http.ResponseWriter, req *http.Request) {

	seelog.Infof("Router Test : %v", req.URL)
	key := mux.Vars(req)["key"]

	reqBody := utils.ReadRequestBody2JSON(req.Body)
	seelog.Debugf("Request Body : %v", string(reqBody))

	reqURL := req.URL.Query()
	seelog.Debugf("Request Params : %v", reqURL)

	var tmplPages []string
	var api string
	switch key {
	case "test":
		tmplPages = append(tmplPages, "views/test/test.html")
	case "runcmd":
		tmplPages = append(tmplPages, "views/test/runcmd.html")
		api = "RunCMD"
	default:
	}

	// tmpl, err := template.ParseFiles(fmt.Sprintf("views/test/%v.html", key))
	tmpl, err := template.ParseFiles(tmplPages...)
	if err != nil {
		seelog.Errorf("template.ParseFiles Error : %v", err)
		return
	}

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

// TestAjaxHandler func(res http.ResponseWriter, req *http.Request)
func TestAjaxHandler(res http.ResponseWriter, req *http.Request) {

	seelog.Infof("Router Test Ajax : %v", req.URL)
	key := mux.Vars(req)["key"]

	reqBody := utils.ReadRequestBody2JSON(req.Body)
	seelog.Debugf("Request Body : %v", string(reqBody))

	reqURL := req.URL.Query()
	seelog.Debugf("Request Params : %v", reqURL)

	var retdata []byte
	switch key {

	default:
	}

	res.Write(retdata)

}
