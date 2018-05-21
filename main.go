package main

import (
	"app/httpsvr"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 5s", func() {
		fmt.Println("Every 5s")
		fmt.Println(time.Now())
		time.Sleep(time.Second * 8)
	})
	c.Start()
	log.Fatalln(httpsvr.StartHTTP())
}
