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
	vars := mux.Vars(req)
	key := vars["key"]

	tmplPages, err := models.GetTmplPages(key)
	if err != nil {
		log.Printf("httpsvr.handler.html.HTMLHandler -> models.GetTmplPages Error : %v\n", err)
		return
	}

	// tmpl, err := template.ParseFiles(fmt.Sprintf("views/test/%v.html", key))
	tmpl, err := template.ParseFiles(tmplPages...)
	if err != nil {
		log.Printf("httpsvr.handler.html.HTMLHandler -> template.ParseFiles Error : %v\n", err)
		return
	}

	api, err := models.GetAPIMap("html", key)
	if err != nil {
		log.Printf("httpsvr.handler.html.HTMLHandler -> models.GetTmplAPI Error : %v\n", err)
		return
	}

	if api != "" {
		fc, err := utils.FuncCall(api)
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
