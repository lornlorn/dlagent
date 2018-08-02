package models

// AjaxReturn struct
type AjaxReturn struct {
	RetCode string
	RetMsg  string
}

// AjaxReturnWithData struct
type AjaxReturnWithData struct {
	Ret     AjaxReturn
	RetData interface{}
}
