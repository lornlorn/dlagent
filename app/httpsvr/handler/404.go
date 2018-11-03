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
	// if req.URL.Path == "/" {
	// 	http.Redirect(res, req, "/index", http.StatusFound)
	// } else if req.URL.Path == "/aaa" {
	// 	for i := 0; i <= 10; i++ {
	// 		go fmt.Printf("%v %v\n", req.URL, i)
	// 		// time.Sleep(10 * time.Second)
	// 	}
	// 	// time.Sleep(10 * time.Second)
	// 	fmt.Fprintln(res, "Route aaa Finish")
	// 	return
	// }
	tmpl, err := template.ParseFiles("views/error/404.html")
	if err != nil {
		log.Printf("httpsvr.handler.404.NotFoundHandler -> template.ParseFiles Error : %v\n", err)
		return
	}
	tmpl.Execute(res, req.URL)
}
