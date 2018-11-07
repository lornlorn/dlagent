package handler

import (
	"app/models"
	"app/utils"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AjaxHandler func(res http.ResponseWriter, req *http.Request)
func AjaxHandler(res http.ResponseWriter, req *http.Request) {

	log.Printf("Route Ajax : %v\n", req.URL)
	vars := mux.Vars(req)
	key := vars["key"]
	// log.Println(key)

	// 获取请求包体(json数据)
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("httpsvr.handler.ajax.AjaxHandler -> ioutil.ReadAll Error : %v\n", err)
		return
	}
	log.Println("Request JSON Content :")
	log.Println(string(reqBody))

	// module := gjson.Get(string(reqBody), "module")
	// shell := gjson.Get(string(reqBody), "data.shell")
	// cmd := gjson.Get(string(reqBody), "data.cmd")
	// log.Println(module, shell, cmd)

	api, err := models.GetSysAPIMap("ajax", key)
	if err != nil {
		log.Printf("httpsvr.handler.ajax.AjaxHandler -> models.GetAjaxAPI Error : %v\n", err)
		return
	}

	if api != "" {
		fc, err := utils.FuncCall(api, reqBody)
		if err != nil {
			log.Printf("httpsvr.handler.ajax.AjaxHandler -> utils.FuncCall Error : %v\n", err)
			return
		}
		log.Println(string(fc[0].Bytes()))

		res.Write(fc[0].Bytes())
	} else {
		res.Write(utils.GetAjaxRetJSON("9999", errors.New("API调用失败")))
	}

}
