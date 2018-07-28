package models

import (
	"app/db"
	"log"
)

/*
TemplatePages struct map table template_pages
*/
type TemplatePages struct {
	Key  string `xorm:"VARCHAR(32) NOT NULL"`
	Page string `xorm:"VARCHAR(128) NOT NULL"`
}

/*
GetTmplPages func(key string) ([]string, error)
Return HTML pages path as []string
*/
func GetTmplPages(key string) ([]string, error) {

	tmplPages := make([]TemplatePages, 0)
	// if err := db.Engine.Where("cron_status = ? and upper(system_enname) like ?", "READY", strings.ToUpper(enkeyword)+"%").Find(&crons); err != nil {
	if err := db.Engine.Where("key = ?", key).Find(&tmplPages); err != nil {
		// return nil, err
		log.Println(err)
		return nil, err
	}

	pages := make([]string, len(tmplPages))
	for i, v := range tmplPages {
		log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
		pages[i] = v.Page
	}

	return pages, nil
}
