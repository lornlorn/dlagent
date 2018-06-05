package monitor

import (
	"app/models"
	"errors"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

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

		if cr.CronType == "default" {
			c.AddFunc(cr.CronExpression, func() {
				log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
			})

			// switch cr.CronCmd {
			// case "cpu":

			// case "mem":

			// default:
			// }

		} else if cr.CronType == "user-defined" {
			c.AddFunc(cr.CronExpression, func() {
				log.Println(cr.CronType, cr.CronExpression, cr.CronCmd)
				// cronargs := strings.Split(cr.CronArgs, " ")
				ret, err := runCmd(cr.CronCmd, cr.CronArgs)
				if err != nil {
					log.Printf("Run Command Fail : %v\n", err)
					// return err
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

func runCmd(croncmd string, cronargs string) ([]byte, error) {
	// 转换为可变长数组
	args := strings.Split(cronargs, " ")
	cmd := exec.Command(croncmd, args...)
	t, _ := exec.LookPath(croncmd)
	log.Println(t)
	log.Println(filepath.Base(croncmd))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	var out = make([]byte, 0, 1024)
	for {
		tmp := make([]byte, 128)
		n, err := stdout.Read(tmp)
		out = append(out, tmp[:n]...)
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	return out, nil
}
