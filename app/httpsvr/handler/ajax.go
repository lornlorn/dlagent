package handler

import (
	"app/models"
	"app/utils"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AjaxHandler func(res http.ResponseWriter, req *http.Request)
func AjaxHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Ajax : %v\n", req.URL)
	key := mux.Vars(req)["key"]

	reqBody := utils.ReadRequestBody2JSON(req.Body)
	log.Println(string(reqBody))

	reqURL := req.URL.Query()

	api, err := models.GetAPI("ajax", key)
	if err != nil {
		log.Printf("httpsvr.handler.ajax.AjaxHandler -> models.GetAPI Error : %v\n", err)
		return
	}

	if api != "" {
		fc, err := utils.FuncCall(api, reqBody, reqURL)
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
