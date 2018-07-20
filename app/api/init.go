package api

import "app/utils"

/*
InitAPIFuncMap func()
*/
func InitAPIFuncMap() {
	var ajax Ajax
	var tmpl Tmpl
	utils.InitFunctionMap(&ajax, &tmpl)
}
