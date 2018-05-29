package monitor

import (
	"app/models"
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
	C = cron.New()
	log.Println("-> Initialize Cron...")
	err := loadJobs(C)
	if err != nil {
		log.Printf("Load Jobs Fail : %v\n", err)
		return err
	}
	log.Println("-> Begin Run Jobs...")
	C.Start()
	return nil
}

func loadJobs(c *cron.Cron) error {
	// c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	// c.AddFunc("@every 5s", func() {
	// 	fmt.Println("Every 5s")
	// 	fmt.Println(time.Now())
	// 	time.Sleep(time.Second * 8)
	// })

	crons, err := models.GetCronList()
	if err != nil {
		log.Printf("Get Crons Fail : %v\n", err)
		return err
	}

	for i, v := range crons {
		log.Printf("DataIndex : %v, DataContent : %v\n", i, v)

		cr := v

		// c.AddFunc(v.CronExpression, runjob)
		switch cr.CronType {
		case "collect":
			c.AddFunc(cr.CronExpression, func() {
				log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
			})
		case "自定义":
			c.AddFunc(cr.CronExpression, func() {
				log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
			})
		default:
			log.Println(cr.CronType)
		}
	}

	return nil
}

func runjob() {
	fmt.Println(time.Now())
}
