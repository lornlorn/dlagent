package api

import "app/utils"

/*
InitAPIFuncMap func()
*/
func InitAPIFuncMap() {
	var ajax Ajax
	var html HTML
	utils.InitFunctionMap(&ajax, &html)
}
