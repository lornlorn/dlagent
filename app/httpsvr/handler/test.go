package handler

import (
	"app/test"
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
	log.Println(key)
	tmpl, err := template.ParseFiles("views/test/test.html")
	if err != nil {
		log.Printf("httpsvr.handler.test.TestHandler -> template.ParseFiles Error : %v\n", err)
		return
	}

	tmpl.Execute(res, nil)
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

	test.GetWFsTest()

	res.Write(reqBody)
}
