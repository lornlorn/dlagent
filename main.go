package main

import (
	"app/db"
	"app/httpsvr"
	"app/monitor"
	"app/utils"
	"log"
)

func main() {
	var err error

	// Read Configuration
	dbtype, _ := utils.ReadConf("db", "dbtype")
	dbstr, _ := utils.ReadConf("db", "dbstr")

	// Init DB
	err = db.InitDB(dbtype, dbstr)
	if err != nil {
		log.Printf("DB Initialize Fail : %v\n", err)
		return
	}
	defer db.Engine.Close()

	// Start Monitor
	err = monitor.Start()
	if err != nil {
		log.Printf("Monitor Start Fail : %v\n", err)
		return
	}

	// Start HTTP Server
	log.Fatalln(httpsvr.StartHTTP())
}
