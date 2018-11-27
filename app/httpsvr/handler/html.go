package handler

import (
	"app/models"
	"app/utils"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HTMLHandler func(res http.ResponseWriter, req *http.Request)
func HTMLHandler(res http.ResponseWriter, req *http.Request) {

	log.Printf("Route HTML : %v\n", req.URL)
	key := mux.Vars(req)["key"]

	reqBody := utils.ReadRequestBody2JSON(req.Body)
	log.Println(string(reqBody))

	reqURL := req.URL.Query()

	tmplPages, err := models.GetTmpls(key)
	if err != nil {
		log.Printf("httpsvr.handler.html.HTMLHandler -> models.GetTmpls Error : %v\n", err)
		return
	}

	// tmpl, err := template.ParseFiles(fmt.Sprintf("views/test/%v.html", key))
	tmpl, err := template.ParseFiles(tmplPages...)
	if err != nil {
		log.Printf("httpsvr.handler.html.HTMLHandler -> template.ParseFiles Error : %v\n", err)
		return
	}

	api, err := models.GetAPI("html", key)
	if err != nil {
		log.Printf("httpsvr.handler.html.HTMLHandler -> models.GetAPI Error : %v\n", err)
		return
	}

	if api != "" {
		fc, err := utils.FuncCall(api, reqBody, reqURL)
		if err != nil {
			log.Printf("httpsvr.handler.html.HTMLHandler -> utils.FuncCall Error : %v\n", err)
			return
		}
		log.Println(fc[0].Interface())

		tmpl.Execute(res, fc[0].Interface())
	} else {
		tmpl.Execute(res, nil)
	}
}
