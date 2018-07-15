package scheduler

import (
	"app/models"
	"log"

	"github.com/robfig/cron"
)

/*
C *cron.Cron
*/
var C *cron.Cron

func startCron() error {
	log.Println("Start Cron ...")
	C = cron.New()
	log.Println("-> Load Cron List ...")
	err := loadJobs(C)
	if err != nil {
		log.Printf("Load Jobs Fail : %v\n", err)
		return err
	}
	log.Println("-> Begin Run Cron Jobs ...")
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

	jobs, err := models.GetJobList()
	if err != nil {
		log.Printf("Get Jobs Fail : %v\n", err)
		return err
	}

	for i, v := range jobs {
		log.Printf("DataIndex : %v, DataContent : %v\n", i, v)

		job := v

		switch job.JobType {
		case "collect":

		case "cron":

		case "plan":

		default:
		}

		/*
			if job.JobType == "default" { // 默认采集待完善
				c.AddFunc(cr.CronExpression, func() {
					// log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
				})

			} else if cr.CronType == "user-defined" { // 自定义调用外部命令
				c.AddFunc(cr.CronExpression, func() {
					// log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
					ret, err := RunCmd(cr.CronSh, cr.CronCmd)
					if err != nil {
						log.Printf("Run Command Fail : %v\n", err)
					}
					log.Println(string(ret))
				})
			} else {
				log.Println("Error CronType")
				return errors.New("Error CronType")
			}
		*/
	}
	return nil
}
