package handler

import (
	"html/template"
	"net/http"

	seelog "github.com/cihub/seelog"
)

// IndexHandler func(res http.ResponseWriter, req *http.Request)
func IndexHandler(res http.ResponseWriter, req *http.Request) {

	seelog.Infof("Router Index : %v", req.URL)

	tmpl, err := template.ParseFiles("views/html/index.html")
	if err != nil {
		seelog.Errorf("template.ParseFiles Error : %v", err)
		return
	}

	tmpl.Execute(res, nil)

}
