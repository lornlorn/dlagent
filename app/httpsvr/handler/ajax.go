package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AjaxHandler func(res http.ResponseWriter, req *http.Request)
func AjaxHandler(res http.ResponseWriter, req *http.Request) {

	log.Printf("Route Ajax : %v\n", req.URL)
	vars := mux.Vars(req)
	key := vars["key"]

	log.Println(key)

	res.Write([]byte(key))

}
