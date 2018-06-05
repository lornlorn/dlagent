package models

import (
	"app/db"
	"log"
)

// Cron struct map db table cron
type Cron struct {
	CronId         int    `xorm:"not null pk unique INTEGER"` // xorm
	CronName       string `xorm:"not null VARCHAR(64)"`
	CronType       string `xorm:"not null VARCHAR(16)"`
	CronExpression string `xorm:"not null VARCHAR(128)"`
	CronCmd        string `xorm:"not null VARCHAR(256)"`
	CronArgs       string `xorm:"VARCHAR(256)"`
	CronStatus     string `xorm:"not null VARCHAR(16)"`
	CronRemark     string `xorm:"VARCHAR(512)"`
}

// Save insert method
func (c *Cron) Save() error {
	// affected, err := db.Engine.Insert(d)
	_, err := db.Engine.Insert(c)
	if err != nil {
		return err
	}
	return nil
}

// GetCronList func() ([]Cron, error)
func GetCronList() ([]Cron, error) {

	crons := make([]Cron, 0)
	// if err := db.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	if err := db.Engine.Where("cron_status = ?", "READY").Find(&crons); err != nil {
		// return nil, err
		log.Println(err)
		return nil, err
	}

	// for i, v := range crons {
	// 	log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return crons, nil
}
