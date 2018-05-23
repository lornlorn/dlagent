package main

import (
	"app/httpsvr"
	"app/monitor"
	"log"
)

func main() {
	monitor.StartCron()
	log.Fatalln(httpsvr.StartHTTP())
}
