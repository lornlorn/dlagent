package handler

import (
	"app/test"
	"app/utils"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TestHandler func(res http.ResponseWriter, req *http.Request)
func TestHandler(res http.ResponseWriter, req *http.Request) {

	log.Printf("Route Test : %v\n", req.URL)
	vars := mux.Vars(req)
	key := vars["key"]

	query := req.URL.Query()
	// log.Println(query["WfiId"][0])
	log.Println(query)

	var tmplPages []string
	var api string
	switch key {
	case "datatables":
		tmplPages = append(tmplPages, "views/test/datatables.html")
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
		log.Printf("httpsvr.handler.test.TestHandler -> template.ParseFiles Error : %v\n", err)
		return
	}

	if api != "" {
		fc, err := utils.FuncCall(api)
		if err != nil {
			log.Printf("httpsvr.handler.test.TestHandler -> utils.FuncCall Error : %v\n", err)
			return
		}
		log.Println(fc[0].Interface())

		tmpl.Execute(res, fc[0].Interface())
	} else {
		tmpl.Execute(res, nil)
	}
}

// TestAjaxHandler func(res http.ResponseWriter, req *http.Request)
func TestAjaxHandler(res http.ResponseWriter, req *http.Request) {

	log.Printf("Route Test Ajax : %v\n", req.URL)
	vars := mux.Vars(req)
	key := vars["key"]

	// 获取请求包体(json数据)
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("httpsvr.handler.test.TestAjaxHandler ioutil.ReadAll Error : %v\n", err)
		return
	}
	log.Println("Request JSON Content :")
	log.Println(string(reqBody))

	// module := gjson.Get(string(reqBody), "module")
	// shell := gjson.Get(string(reqBody), "data.shell")
	// cmd := gjson.Get(string(reqBody), "data.cmd")
	// log.Println(module, shell, cmd)

	var retdata []byte
	switch key {
	case "datatables":
		retdata = test.GetWFsTest()
	case "delete":
		retdata = test.DelWfByID(reqBody)
	default:
		retdata = test.GetWFsTest()
	}

	res.Write(retdata)
}
