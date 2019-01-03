package models

import (
	"app/utils"
	"errors"

	seelog "github.com/cihub/seelog"
)

/*
TbComponent struct map to table tb_component
*/
type TbComponent struct {
	CompId     int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	CompName   string `xorm:"VARCHAR(128) NOT NULL"`
	CompStatus string `xorm:"VARCHAR(8) NOT NULL"`
	CompCmd    string `xorm:"VARCHAR(1024) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(19)"`
	ModifyTime string `xorm:"VARCHAR(19)"`
}

/*
NewComponent struct map to table tb_component without column CompId
*/
type NewComponent struct {
	CompName   string `xorm:"VARCHAR(128) NOT NULL"`
	CompStatus string `xorm:"VARCHAR(8) NOT NULL"`
	CompCmd    string `xorm:"VARCHAR(1024) NOT NULL"`
	CreateTime string `xorm:"VARCHAR(19)"`
	ModifyTime string `xorm:"VARCHAR(19)"`
}

/*
TableName xorm mapper
NewComponent struct map to table tb_component
*/
func (comp NewComponent) TableName() string {
	return "tb_component"
}

// Save insert method
func (comp NewComponent) Save() error {
	affected, err := utils.Engine.Insert(comp)
	if err != nil {
		// seelog.Errorf("utils.Engine.Insert Error : %v", err)
		return err
	}
	seelog.Debugf("%v insert : %v", affected, comp)

	return nil
}

// Update method
func (comp TbComponent) Update() error {
	affected, err := utils.Engine.ID(comp.CompId).Update(comp)
	if err != nil {
		// seelog.Errorf("utils.Engine.ID.Update Error : %v", err)
		return err
	}
	seelog.Debugf("%v update : %v", affected, comp)

	return nil
}

/*
GetComponents func() ([]TbComponent, error)
*/
func GetComponents() ([]TbComponent, error) {
	components := make([]TbComponent, 0)

	if err := utils.Engine.Find(&components); err != nil {
		// seelog.Errorf("utils.Engine.Find Error : %v", err)
		return nil, err
	}

	return components, nil
}

/*
GetComponentByID func(compid int) (TbComponent, error)
*/
func GetComponentByID(compid int) (TbComponent, error) {
	comp := new(TbComponent)
	comp.CompId = compid

	has, err := utils.Engine.Get(comp)
	if err != nil {
		// seelog.Errorf("utils.Engine.Get Error : %v", err)
		return TbComponent{}, err
	}

	if !has {
		// seelog.Debug("Get 0 row")
		return TbComponent{}, errors.New("Get 0 rows")
	}

	return *comp, nil
}

/*
关联关系改变，临时作废

GetWorkflowDtlByWfiID func(wfiid int) ([]SysWorkflowDtl, error)

func GetWorkflowDtlByWfiID(compid int) ([]TbComponent, error) {

	details := make([]SysWorkflowDtl, 0)

	if err := utils.Engine.Where("wfi_id = ?", wfiid).Asc("wfd_seq").Find(&details); err != nil {
		seelog.Errorf("utils.Engine.Where Error : %v", err)
		return nil, err
	}

	return details, nil

}
*/

/*
DelComponentByID func(compid int) error
*/
func DelComponentByID(compid int) error {
	comp := new(TbComponent)
	comp.CompId = compid

	affected, err := utils.Engine.Delete(comp)
	if err != nil {
		// seelog.Errorf("utils.Engine.Delete Error : %v", err)
		return err
	}
	seelog.Debugf("%v delete : %v", affected, comp)

	return nil
}
