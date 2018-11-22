package main

import (
	"app/api"
	"app/httpsvr"
	"app/utils"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe("localhost:8765", nil))
	}()

	var err error

	// Read Configuration
	log.Println("Load Configuration ...")
	dbtype, _ := utils.ReadConf("db", "dbtype")
	dbstr, _ := utils.ReadConf("db", "dbstr")

	// Init Reflect Functions
	log.Println("Initialize Reflect Function Map ...")
	api.InitAPIFuncMap()

	// Init DB
	log.Println("Initialize Database Connect ...")
	err = utils.InitDB(dbtype, dbstr)
	if err != nil {
		log.Printf("DB Initialize Fail : %v\n", err)
		return
	}
	defer utils.Engine.Close()

	/*
		// Start Monitor
		log.Println("Scheduler Start ...")
		err = scheduler.Start()
		if err != nil {
			log.Printf("Scheduler Start Fail : %v\n", err)
			return
		}
	*/

	// Start HTTP Server
	log.Println("Start HTTP Server ...")
	log.Fatalln(httpsvr.StartHTTP())
}
