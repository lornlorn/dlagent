package monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

/*
C *cron.Cron
*/
var C *cron.Cron

func startCron() error {
	log.Println("Start Monitor...")
	log.Println("-> Initialize Cron...")
	C = cron.New()
	err := loadJobs(C)
	if err != nil {
		log.Printf("Load Jobs Fail : %v\n", err)
		return err
	}
	C.Start()
	return nil
}

func loadJobs(c *cron.Cron) error {
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 5s", func() {
		fmt.Println("Every 5s")
		fmt.Println(time.Now())
		time.Sleep(time.Second * 8)
	})
	return nil
}
