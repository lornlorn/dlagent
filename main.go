package main

import (
	"app/httpsvr"
	"app/monitor"
	"log"
	"time"
)

func main() {
	err := monitor.Start()
	if err != nil {
		log.Printf("Monitor Start Fail : %v\n", err)
		return
	}
	log.Println("Cron start")
	time.Sleep(time.Second * 5)
	monitor.Stop()
	log.Println("Cron stop")
	time.Sleep(time.Second * 5)
	monitor.Start()
	log.Println("Cron start")
	time.Sleep(time.Second * 5)
	log.Fatalln(httpsvr.StartHTTP())
}
