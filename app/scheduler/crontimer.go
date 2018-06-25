package scheduler

import (
	"app/models"
	"errors"
	"log"

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
	//  fmt.Println("Every 5s")
	//  fmt.Println(time.Now())
	//  time.Sleep(time.Second * 8)
	// })

	crons, err := models.GetCronList()
	if err != nil {
		log.Printf("Get Crons Fail : %v\n", err)
		return err
	}

	for i, v := range crons {
		log.Printf("DataIndex : %v, DataContent : %v\n", i, v)

		cr := v

		if cr.CronType == "default" { // 默认采集待完善
			c.AddFunc(cr.CronExpression, func() {
				// log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
			})

			// switch cr.CronCmd {
			// case "cpu":

			// case "mem":

			// default:
			// }

		} else if cr.CronType == "user-defined" { // 自定义调用外部命令
			c.AddFunc(cr.CronExpression, func() {
				// log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
				ret, err := runCmd(cr.CronSh, cr.CronCmd)
				if err != nil {
					log.Printf("Run Command Fail : %v\n", err)
				}
				log.Println(string(ret))
			})
		} else {
			log.Println("Error CronType")
			return errors.New("Error CronType")
		}
	}
	return nil
}
