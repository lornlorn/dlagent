package handler

import (
	"html/template"
	"net/http"

	seelog "github.com/cihub/seelog"
)

// NotFoundHandler func(res http.ResponseWriter, req *http.Request)
/*
Route Not Found 404 Page
And
Route "/" Direct To "/index"
*/
func NotFoundHandler(res http.ResponseWriter, req *http.Request) {
	seelog.Infof("Router 404 : %v", req.URL)

	if req.URL.Path == "/favicon.ico" {
		seelog.Debug("Request A favicon")
		http.ServeFile(res, req, "./static/img/favicon.ico")
		return
	}

	tmpl, err := template.ParseFiles("./views/error/404.html")
	if err != nil {
		seelog.Errorf("template.ParseFiles Error : %v", err)
		return
	}
	tmpl.Execute(res, req.URL)
}
