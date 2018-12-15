package main

import (
	"app/api"
	"app/httpsvr"
	"app/utils"
	"log"
	"net/http"
	_ "net/http/pprof"

	seelog "github.com/cihub/seelog"
)

const (
	SEELOG_CFG = "./config/seelog.xml" // SEELOG_CFG seelog config file path
	APP_CFG    = "./config/app.conf"   // APP_CFG app config file path
)

func main() {
	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe("localhost:8765", nil))
	}()

	var err error
	var msg string

	// Initialize Logger
	msg = "1 -> Initialize Logger"
	err = utils.InitLogger(SEELOG_CFG)
	if err != nil {
		log.Fatalf("%v Error : %v", msg, err)
		panic("Exit!")
		// return
	}
	seelog.Infof("%v Success !", msg)

	// Read Configuration
	msg = "2 -> Load Configuration"
	err = utils.InitConfig(APP_CFG)
	if err != nil {
		seelog.Criticalf("%v Error : %v", msg, err)
		panic("Exit!")
	}
	seelog.Infof("%v Success !", msg)

	// Init Reflect Functions
	msg = "3 -> Initialize Reflect Function Map"
	api.InitAPIFuncMap()
	seelog.Infof("%v Success !", msg)

	// Init DB
	msg = "4 -> Connect Database"
	dbtype := utils.GetConfig("db", "dbtype")
	dbstr := utils.GetConfig("db", "dbstr")
	err = utils.InitDB(dbtype, dbstr)
	if err != nil {
		seelog.Criticalf("%v Error : %v", msg, err)
		panic("Exit!")
	}
	defer utils.Engine.Close()
	seelog.Infof("%v Success !", msg)

	/*
		// Start Monitor
		log.Println("Scheduler Start ...")
		err = scheduler.Start()
		if err != nil {
			log.Printf("Scheduler Start Fail : %v", err)
			return
		}
	*/

	// Start HTTP Server
	msg = "5 -> Starting HTTP Server"
	seelog.Infof("%v !", msg)
	seelog.Info("***Everything is OK !***")
	// log.Fatalln(httpsvr.StartHTTP())
	err = httpsvr.StartHTTP()
	if err != nil {
		seelog.Criticalf("%v Error : %v", msg, err)
		panic("Exit!")
	}
	seelog.Infof("%v Success !", msg)
}
