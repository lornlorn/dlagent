package handler

import (
	"app/scheduler"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TestHandler func(res http.ResponseWriter, req *http.Request)
func TestHandler(res http.ResponseWriter, req *http.Request) {
	scheduler.Stop()

	log.Printf("Route Test : %v\n", req.URL)
	vars := mux.Vars(req)
	subroute := vars["page"]
	tmpl, err := template.ParseFiles(fmt.Sprintf("views/test/%v.html", subroute))
	if err != nil {
		log.Printf("Parse Error : %v\n", err)
		return
	}
	tmpl.Execute(res, nil)

}
