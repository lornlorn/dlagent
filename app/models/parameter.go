package models

import (
	"app/utils"
	"errors"

	seelog "github.com/cihub/seelog"
)

/*
TbParameter struct map to table tb_parameter
*/
type TbParameter struct {
	ParamId      int    `xorm:"INTEGER NOT NULL UNIQUE PK"`
	CompId       int    `xorm:"INTEGER NOT NULL"`
	ParamSeq     int    `xorm:"INTEGER NOT NULL"`
	ParamName    string `xorm:"VARCHAR(128) NOT NULL"`
	ParamDefault string `xorm:"VARCHAR(128)"`
	CreateTime   string `xorm:"VARCHAR(15)"`
	ModifyTime   string `xorm:"VARCHAR(15)"`
}

/*
NewParameter struct map to table tb_parameter without column ParamId
*/
type NewParameter struct {
	CompId       int    `xorm:"INTEGER NOT NULL"`
	ParamSeq     int    `xorm:"INTEGER NOT NULL"`
	ParamName    string `xorm:"VARCHAR(128) NOT NULL"`
	ParamDefault string `xorm:"VARCHAR(128)"`
	CreateTime   string `xorm:"VARCHAR(15)"`
	ModifyTime   string `xorm:"VARCHAR(15)"`
}

/*
TableName xorm mapper
NewParameter struct map to table tb_parameter
*/
func (param NewParameter) TableName() string {
	return "tb_parameter"
}

// Save insert method
func (param NewParameter) Save() error {
	affected, err := utils.Engine.Insert(param)
	if err != nil {
		// seelog.Errorf("utils.Engine.Insert Error : %v\n", err)
		return err
	}
	seelog.Debugf("%v insert : %v", affected, param)

	return nil
}

// Update method
func (param TbParameter) Update() error {
	affected, err := utils.Engine.ID(param.ParamId).Update(param)
	if err != nil {
		// seelog.Errorf("utils.Engine.ID.Update Error : %v", err)
		return err
	}
	seelog.Debugf("%v update : %v", affected, param)

	return nil
}

/*
GetParametersByCompID func(compid int) ([]TbParameter, error)
*/
func GetParametersByCompID(compid int) ([]TbParameter, error) {
	params := make([]TbParameter, 0)

	if err := utils.Engine.Where("comp_id = ?", compid).Asc("param_seq").Find(&params); err != nil {
		// seelog.Errorf("utils.Engine.Where Error : %v", err)
		return nil, err
	}

	return params, nil
}

/*
DelParameterByID func(paramid int) error
*/
func DelParameterByID(paramid int) error {
	param := new(TbParameter)
	param.ParamId = paramid

	affected, err := utils.Engine.Delete(param)
	if err != nil {
		// seelog.Errorf("utils.Engine.Delete Error : %v", err)
		return err
	}
	seelog.Debugf("%v delete : %v", affected, param)

	return nil
}

/*
GetLastParameterByCompID func(compid int) (TbParameter, error)
*/
func GetLastParameterByCompID(compid int) (TbParameter, error) {
	param := new(TbParameter)

	has, err := utils.Engine.Where("comp_id = ?", compid).Desc("param_seq").Get(param)
	if err != nil {
		// seelog.Errorf("utils.Engine.Desc.Get Error : %v", err)
		return TbParameter{}, err
	}

	if has {
		// seelog.Debugf("Workflow Last Param : %v", wfp)
		return *param, nil
	}
	// seelog.Debugf("This Dtl [%v] Has No Parameters", wfdid)

	return TbParameter{}, errors.New("No Records")
}
