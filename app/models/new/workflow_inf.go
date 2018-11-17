package new

import (
	"app/db"
	"log"
)

/*
SysWorkflowInf struct map to table sys_workflow_inf
*/
type SysWorkflowInf struct {
	WfiId      int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	WfiName    string `xorm:"VARCHAR(128) NOT NULL"`
	WfiDesc    string `xorm:"VARCHAR(1024)"`
	WfiStatus  string `xorm:"VARCHAR(8) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(15)"`
	ModifyTime string `xorm:"VARCHAR(15)"`
}

/*
GetWorkflows func() ([]SysWorkflowInf, error)
*/
func GetWorkflows() ([]SysWorkflowInf, error) {

	workflows := make([]SysWorkflowInf, 0)

	// if err := db.Engine.Where("job_type = ?", jobType).Find(&jobs); err != nil {
	if err := db.Engine.Find(&workflows); err != nil {
		// return nil, err
		log.Println(err)
		return nil, err
	}

	// for i, v := range crons {
	//  log.Printf("DataIndex : %v, DataContent : %v\n", i, v)
	// }

	return workflows, nil
}
