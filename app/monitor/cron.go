package monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

/*
StartCron func()
*/
func StartCron() error {
	log.Println("Start Monitor...")
	log.Println("-> Initialize Cron...")
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 5s", func() {
		fmt.Println("Every 5s")
		fmt.Println(time.Now())
		time.Sleep(time.Second * 8)
	})
	c.Start()
	return nil
}
