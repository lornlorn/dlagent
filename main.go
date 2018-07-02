package main

import (
	"app/db"
	"app/httpsvr"
	"app/scheduler"
	"app/utils"
	"log"
)

func main() {
	var err error

	// Read Configuration
	log.Println("Load Configuration ...")
	dbtype, _ := utils.ReadConf("db", "dbtype")
	dbstr, _ := utils.ReadConf("db", "dbstr")

	// Init Reflect Functions
	log.Println("Initialize Reflect Function Map ...")
	utils.InitFunctionMap()

	// Init DB
	log.Println("Initialize Database Connect ...")
	err = db.InitDB(dbtype, dbstr)
	if err != nil {
		log.Printf("DB Initialize Fail : %v\n", err)
		return
	}
	defer db.Engine.Close()

	// Start Monitor
	log.Println("Scheduler Start ...")
	err = scheduler.Start()
	if err != nil {
		log.Printf("Scheduler Start Fail : %v\n", err)
		return
	}

	// Start HTTP Server
	log.Println("Start HTTP Server ...")
	log.Fatalln(httpsvr.StartHTTP())
}
