package models

import (
	seelog "github.com/cihub/seelog"
)

/*
ComponentWithParams struct
component with parameters
*/
type ComponentWithParams struct {
	Comp   TbComponent
	Params []TbParameter
}

/*
GetComponentWithParamsByCompID func(compid int) (ComponentWithParams, error)
*/
func GetComponentWithParamsByCompID(compid int) (ComponentWithParams, error) {
	comp, err := GetComponentByID(compid)
	if err != nil {
		seelog.Errorf("GetComponentByID Error : %v", err)
		return ComponentWithParams{}, err
	}
	params, err := GetParametersByCompID(compid)
	if err != nil {
		seelog.Errorf("GetParametersByCompID Error : %v", err)
		return ComponentWithParams{}, err
	}

	var compWithParams ComponentWithParams
	compWithParams.Comp = comp
	compWithParams.Params = params

	seelog.Debugf("ComponentWithParams : %v", compWithParams)

	return compWithParams, nil
}
