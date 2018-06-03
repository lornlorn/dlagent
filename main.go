package main

import (
	"app/db"
	"app/httpsvr"
	"app/monitor"
	"log"
)

func main() {
	var err error

	// Init DB
	err = db.InitDB()
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
