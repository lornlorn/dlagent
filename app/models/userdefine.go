package models

// AjaxReturn struct
type AjaxReturn struct {
	RetCode string `json:"retcode"`
	RetMsg  string `json:"retmsg"`
}

// AjaxReturnWithData struct
type AjaxReturnWithData struct {
	RetCode string      `json:"retcode"`
	RetMsg  string      `json:"retmsg"`
	RetData interface{} `json:"retdata"`
}
