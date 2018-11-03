package handler

import (
	"html/template"
	"log"
	"net/http"
)

// IndexHandler func(res http.ResponseWriter, req *http.Request)
func IndexHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route Index : %v\n", req.URL)
	tmpl, err := template.ParseFiles("views/html/index.html")
	if err != nil {
		log.Printf("httpsvr.handler.index.IndexHandler -> template.ParseFiles Error : %v\n", err)
		return
	}

	tmpl.Execute(res, nil)
}
