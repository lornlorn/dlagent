package httpsvr

import (
	"app/httpsvr/handler"
	"net/http"
	"time"

	seelog "github.com/cihub/seelog"
	"github.com/gorilla/mux"
)

/*
StartHTTP func()
*/
func StartHTTP() error {

	r := mux.NewRouter()
	initRoutes(r)
	seelog.Info("Initialize HTTP Routers Success !")

	svr := &http.Server{
		Handler:      r,
		Addr:         ":5678",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	seelog.Info("Listen HTTP Port And Serve")
	// err := http.ListenAndServe(":8888", r)
	err := svr.ListenAndServe()

	return err

}

func initRoutes(r *mux.Router) {

	// normal router
	r.HandleFunc("/index", handler.IndexHandler)

	// html router
	h := r.PathPrefix("/html").Subrouter()
	h.HandleFunc("/", handler.NotFoundHandler)
	// h.HandleFunc("/{key}", handler.NotFoundHandler)
	// h.HandleFunc("/{group}/{module}", handler.HTMLHandler)
	h.HandleFunc("/{key}", handler.HTMLHandler)

	// ajax router
	a := r.PathPrefix("/ajax").Subrouter()
	a.HandleFunc("/", handler.NotFoundHandler)
	a.HandleFunc("/{key}", handler.AjaxHandler)

	// test router
	t := r.PathPrefix("/test").Subrouter()
	t.HandleFunc("/", handler.NotFoundHandler)
	t.HandleFunc("/{key}", handler.TestHandler)
	t.HandleFunc("/ajax/{key}", handler.TestAjaxHandler)

	// static resource router
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// other root subrouter
	root := r.PathPrefix("/").Subrouter()
	root.HandleFunc("/", handler.IndexHandler)
	// root.Handle("/favicon.ico", http.StripPrefix("/favicon.ico", http.FileServer(http.Dir("static/img/favicon.ico"))))
	// root.Handle("/favicon.ico", http.FileServer(http.Dir("static/img/favicon.ico")))
	root.HandleFunc("/{key}", handler.NotFoundHandler)

	// http.HandleFunc("/", notFoundHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}
