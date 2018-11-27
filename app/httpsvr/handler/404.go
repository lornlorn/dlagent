package handler

import (
	"html/template"
	"log"
	"net/http"
)

// NotFoundHandler func(res http.ResponseWriter, req *http.Request)
/*
Route Not Found 404 Page
And
Route "/" Direct To "/index"
*/
func NotFoundHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Route 404 : %v\n", req.URL)

	if req.URL.Path == "/favicon.ico" {
		http.ServeFile(res, req, "./static/img/favicon.ico")
		return
	}

	tmpl, err := template.ParseFiles("views/error/404.html")
	if err != nil {
		log.Printf("httpsvr.handler.404.NotFoundHandler -> template.ParseFiles Error : %v\n", err)
		return
	}
	tmpl.Execute(res, req.URL)
}
