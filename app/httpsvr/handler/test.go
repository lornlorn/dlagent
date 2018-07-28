package handler

import (
	"app/models"
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

	tmplPages, err := models.GetTmplPages(key)
	if err != nil {
		log.Printf("httpsvr.handler.test.TestHandler models.GetTmplPages Error : %v\n", err)
		return
	}

	// tmpl, err := template.ParseFiles(fmt.Sprintf("views/test/%v.html", key))
	tmpl, err := template.ParseFiles(tmplPages...)
	if err != nil {
		log.Printf("httpsvr.handler.test.TestHandler template.ParseFiles Error : %v\n", err)
		return
	}

	switch key {
	case "joblist":
		fc, err := utils.FuncCall("GetJobList")
		if err != nil {
			log.Printf("httpsvr.handler.test.TestHandler utils.FuncCall Error : %v\n", err)
			return
		}
		log.Println(fc[0].Interface())

		tmpl.Execute(res, fc[0].Interface())
	default:
		tmpl.Execute(res, nil)
	}
}

// TestAjaxHandler func(res http.ResponseWriter, req *http.Request)
func TestAjaxHandler(res http.ResponseWriter, req *http.Request) {

	log.Printf("Route Test Ajax : %v\n", req.URL)
	vars := mux.Vars(req)
	key := vars["key"]
	log.Println(key)

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

	fc, err := utils.FuncCall(key, reqBody)
	if err != nil {
		log.Printf("httpsvr.handler.test.TestAjaxHandler utils.FuncCall Error : %v\n", err)
		return
	}
	log.Println(string(fc[0].Bytes()))

	res.Write(fc[0].Bytes())
}
